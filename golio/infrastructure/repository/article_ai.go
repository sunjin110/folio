package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"sort"
	"sync"

	"github.com/sashabaranov/go-openai"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/chatgpt"
	cdto "github.com/sunjin110/folio/golio/infrastructure/chatgpt/dto"
)

type articleAI struct {
	chatGPTClient          chatgpt.Client
	googleCustomSearchRepo repository.GoogleCustomSearch
	htmlContentRepo        repository.HtmlContent
}

func NewArticleAI(chatGPTClient chatgpt.Client,
	googleCustomSearch repository.GoogleCustomSearch, htmlContentRepo repository.HtmlContent) repository.ArticleAI {
	return &articleAI{
		chatGPTClient:          chatGPTClient,
		googleCustomSearchRepo: googleCustomSearch,
		htmlContentRepo:        htmlContentRepo,
	}
}

func (a *articleAI) ChangeBodyByAI(ctx context.Context, article *model.Article, orderToAI string) (*model.Article, error) {
	output, err := a.chatGPTClient.CreateChatCompletions(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant who helps edit articles based on user instructions.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("I wrote an article or a note. Here's the paragraph: '%s'", article.Body),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: orderToAI,
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed chatGPTClient.CreateChatCompletions. err: %w", err)
	}

	body := output.Choices[0].Message.Content
	article.Body = body
	return article, nil
}

func (a *articleAI) GenerateBodyByAI(ctx context.Context, prompt string) (string, error) {
	output, err := a.generateBodyByAI(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("failed chatGPTClient.CreateChatCompletions. err: %w", err)
	}

	if len(output.Choices) == 0 {
		return "", fmt.Errorf("empty result")
	}

	// tool callsではない場合はそのまま返す
	if output.Choices[0].FinishReason != "tool_calls" {
		return chatgpt.GetMessageContent(output), nil
	}

	// google検索
	// TODO これだと不十分だから、サマライズのAPIを利用して、中身を取るのがいいかも
	arguments := output.Choices[0].Message.ToolCalls[0].Function.Arguments
	htmlContentsJSON, err := a.searchByGoogleForGenerateBody(ctx, arguments)
	if err != nil {
		return "", fmt.Errorf("failed a.searchByGoogleForGenerateBody. arguments: %s, err: %w", arguments, err)
	}

	// Googleの結果を利用して回答を出す
	toolCallID := output.Choices[0].Message.ToolCalls[0].ID
	outputWithGoogle, err := a.generateBodyByAIAndGoogleResult(ctx, toolCallID, arguments, htmlContentsJSON)
	if err != nil {
		return "", fmt.Errorf("failed a.generateBodyByAIAndGoogleResult. err: %w", err)
	}

	return chatgpt.GetMessageContent(outputWithGoogle), nil
}

func (a *articleAI) generateBodyByAI(ctx context.Context, prompt string) (openai.ChatCompletionResponse, error) {
	output, err := a.chatGPTClient.CreateChatCompletions(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a highly knowledgeable assistant.",
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Write detailed articles in Markdown format.",
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Avoid using code blocks in your Markdown responses.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Tools: []openai.Tool{
			{
				Type: openai.ToolTypeFunction,
				Function: &openai.FunctionDefinition{
					Name:        "searchGoogle",
					Description: "When you need information that you don't know or when you need the latest information, search the internet (Google) for it.",
					Parameters: cdto.M{
						"type": "object",
						"properties": cdto.M{
							"keyword": cdto.M{
								"type": "string",
							},
						},
						"required": []string{"keyword"},
					},
				},
			},
		},
		ToolChoice: "auto",
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, fmt.Errorf("failed chatGPTClient.CreateChatCompletions. err: %w", err)
	}
	return output, nil
}

func (a *articleAI) searchByGoogleForGenerateBody(ctx context.Context, arguments string) (string, error) {
	argmentsMap := map[string]string{}
	if err := json.Unmarshal([]byte(arguments), &argmentsMap); err != nil {
		return "", fmt.Errorf("failed json.Unmarshal. err: %w", err)
	}
	googleSearchResults, err := a.googleCustomSearchRepo.Search(ctx, argmentsMap["keyword"])
	if err != nil {
		return "", fmt.Errorf("failed google search. err: %w", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(googleSearchResults))
	mu := sync.Mutex{}
	htmlContents := []*model.HtmlContent{}

	orderMap := map[string]int{} // 関連度が高いやつを優先するようにする

	for i, googleSearchResult := range googleSearchResults {
		googleSearchResult := googleSearchResult

		// orderMap[i] = googleSearchResult.URL
		orderMap[googleSearchResult.URL] = i

		go func() {
			defer wg.Done()
			body, err := a.htmlContentRepo.Get(ctx, googleSearchResult.URL)
			if err != nil {
				slog.Warn("failed get htmlContent", "url", googleSearchResult.URL)
				return
			}

			bodyBytes, err := io.ReadAll(body)
			if err != nil {
				slog.Warn("failed read all body", "url", googleSearchResult.URL)
				return
			}

			baseURL, err := googleSearchResult.GetBaseURL()
			if err != nil {
				slog.Warn("failed get base url", "url", googleSearchResult.URL, "err", err)
				return
			}

			extractor := model.NewHtmlExtractor(string(bodyBytes), baseURL)

			bodyText, err := extractor.ExtractText(ctx)
			if err != nil {
				slog.Warn("failed extract text", "body", body, "url", googleSearchResult.URL)
				return
			}

			if bodyText == "" {
				return
			}

			// TODO 内容の要約APIを将来的に導入して文字数を削減する

			mu.Lock()
			htmlContents = append(htmlContents, &model.HtmlContent{
				Title:    googleSearchResult.Title,
				URL:      googleSearchResult.URL,
				Overview: googleSearchResult.Overview,
				BodyText: bodyText,
			})
			mu.Unlock()
		}()
	}

	wg.Wait()

	// 優先度
	sort.Slice(htmlContents, func(i, j int) bool {
		return orderMap[htmlContents[i].URL] < orderMap[htmlContents[j].URL]
	})

	// 今の所1つにする... token数が圧倒的に足りない...
	htmlContents = htmlContents[0:min(len(htmlContents), 1)]

	htmlContentsJSON, err := json.Marshal(htmlContents)
	if err != nil {
		return "", fmt.Errorf("failed json.Unmarshal. err: %w", err)
	}
	return string(htmlContentsJSON), nil
}

func (a *articleAI) generateBodyByAIAndGoogleResult(ctx context.Context, toolCallID string, arguments string, htmlContentJSON string) (openai.ChatCompletionResponse, error) {
	outputWithGoogle, err := a.chatGPTClient.CreateChatCompletions(ctx,
		openai.ChatCompletionRequest{
			// Model: openai.GPT4,
			Model: openai.GPT4o, // token数の関係でこれにした
			// Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "tool_calls",
					ToolCalls: []openai.ToolCall{
						{
							ID:   toolCallID,
							Type: openai.ToolTypeFunction,
							Function: openai.FunctionCall{
								Name:      "searchGoogle",
								Arguments: arguments,
							},
						},
					},
				},
				{
					Role:       openai.ChatMessageRoleTool,
					Content:    htmlContentJSON,
					ToolCallID: toolCallID,
				},
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Please provide the user with the information obtained from a Google search and the reference URLs.",
				},
			},
		},
	)
	if err != nil {
		return openai.ChatCompletionResponse{}, fmt.Errorf("failed chatGPTClient.CreateChatCompletions. err: %w", err)
	}
	return outputWithGoogle, nil
}
