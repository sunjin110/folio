package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	_ "embed"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/chatgpt"
	chatgptDto "github.com/sunjin110/folio/golio/infrastructure/chatgpt/dto"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
)

var (
	//go:embed query/postgresql/upsert_article_summary.sql
	upsertArticleSummaryPostgreSQL string

	//go:embed query/postgresql/upsert_article_body.sql
	upsertArticleBodyPostgreSQL string
)

type articleV2 struct {
	db                     *sqlx.DB
	chatGPTClient          chatgpt.Client
	googleCustomSearchRepo repository.GoogleCustomSearch
}

func NewArticleV2(ctx context.Context, db *sqlx.DB, chatGPTClient chatgpt.Client,
	googleCustomSearch repository.GoogleCustomSearch) repository.Article {
	return &articleV2{
		db:                     db,
		chatGPTClient:          chatGPTClient,
		googleCustomSearchRepo: googleCustomSearch,
	}
}

func (a *articleV2) CountTotal(ctx context.Context, search *repository.ArticleSearch) (int32, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(*)").From("article_summaries")

	if search != nil && search.Title != nil {
		// TODO indexが利用できないため将来的にPGroongaを利用する
		sb.Where(sb.Like("title", "%"+*search.Title+"%"))
	}

	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := a.db.QueryxContext(ctx, sql, args...)
	if err != nil {
		return -1, fmt.Errorf("failed QueryContext. sql: %s, err: %w", sql, err)
	}

	if !rows.Next() {
		return -1, fmt.Errorf("not found rows. sql: %s", sql)
	}

	var totalCount int32
	if err := rows.Scan(&totalCount); err != nil {
		return -1, fmt.Errorf("failed scan. sql: %s, err: %w", sql, err)
	}
	return totalCount, nil
}

func (a *articleV2) Delete(ctx context.Context, id string) (err error) {
	tx, err := a.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed begin. err: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				slog.ErrorContext(ctx, "failed article delete rollback", "err", err, "id", id)
			}
		}
	}()

	if _, err := tx.ExecContext(ctx, `delete from article_bodies where id = $1`, id); err != nil {
		return fmt.Errorf("failed delete from article_bodies. id: %s, err: %w", id, err)
	}

	if _, err := tx.ExecContext(ctx, "delete from article_summaries where id = $1", id); err != nil {
		return fmt.Errorf("failed delete from article_summaries. id: %s, err: %w", id, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed delete commit. err: %w", err)
	}
	return nil
}

func (a *articleV2) FindSummary(ctx context.Context, sortType repository.SortType, paging *repository.Paging, search *repository.ArticleSearch) ([]*model.ArticleSummary, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("article_summaries")

	if search != nil && search.Title != nil {
		// TODO indexが利用できないため将来的にPGroongaを利用する
		sb.Where(sb.Like("title", "%"+*search.Title+"%"))
	}

	sb.Limit(paging.Limit).Offset(paging.Offset).OrderBy("created_at")
	if sortType == repository.SortTypeAsc {
		sb.Asc()
	} else {
		sb.Desc()
	}
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	summaries := postgres_dto.ArticleSummaries{}
	if err := a.db.SelectContext(ctx, &summaries, sql, args...); err != nil {
		return nil, fmt.Errorf("failed get select article_summaries. sql: %s, err: %w", sql, err)
	}

	return summaries.ToSummariesModel(), nil
}

func (a *articleV2) Get(ctx context.Context, id string) (*model.Article, error) {
	body := &postgres_dto.ArticleBody{}
	if err := a.db.GetContext(ctx, body, "select * from article_bodies where id = $1", id); err != nil {
		return nil, fmt.Errorf("failed get article_bodies. err: %w", err)
	}

	summary := &postgres_dto.ArticleSummary{}
	if err := a.db.GetContext(ctx, summary, "select * from article_summaries where id = $1", id); err != nil {
		return nil, fmt.Errorf("failed get article_summaries. err: %w", err)
	}

	return &model.Article{
		ID:        summary.ID,
		Title:     summary.Title,
		Body:      body.Body,
		Writer:    "",
		CreatedAt: summary.CreatedAt,
		UpdatedAt: summary.UpdatedAt,
	}, nil
}

func (a *articleV2) Insert(ctx context.Context, article *model.Article) error {
	return a.upsert(ctx, article)
}

func (a *articleV2) Update(ctx context.Context, article *model.Article) error {
	return a.upsert(ctx, article)
}

func (a *articleV2) upsert(ctx context.Context, article *model.Article) (err error) {

	tx, err := a.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed upsert transaction begin. err: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				slog.ErrorContext(ctx, "failed article upsert rollback", "err", err)
			}
		}
	}()

	if _, err := tx.NamedExecContext(ctx, upsertArticleBodyPostgreSQL, &postgres_dto.ArticleBody{
		ID:                 article.ID,
		ArticleSummariesID: article.ID,
		Body:               article.Body,
		CreatedAt:          article.CreatedAt,
		UpdatedAt:          article.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed upsert article_bodies. err: %w", err)
	}

	if _, err := tx.NamedExecContext(ctx, upsertArticleSummaryPostgreSQL, &postgres_dto.ArticleSummary{
		ID:        article.ID,
		Title:     article.Title,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed upsert article_summaries. err: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed commit. err: %w", err)
	}
	return nil
}

func (a *articleV2) ChangeBodyByAI(ctx context.Context, article *model.Article, orderToAI string) (*model.Article, error) {
	output, err := a.chatGPTClient.CreateChatCompletions(ctx, &chatgptDto.ChatCompletionsInput{
		Model: chatgpt.GPT4Model,
		Messages: []chatgptDto.Message{
			&chatgptDto.SystemMessage{
				Role:    "system",
				Content: "You are a helpful assistant who helps edit articles based on user instructions.",
			},
			&chatgptDto.UserMessage{
				Role:    "user",
				Content: fmt.Sprintf("I wrote an article or a note. Here's the paragraph: '%s'", article.Body),
			},
			&chatgptDto.UserMessage{
				Role:    "user",
				Content: orderToAI,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed chatGPTClient.CreateChatCompletions. err: %w", err)
	}

	body := output.GetMessage()

	article.Body = body
	return article, nil
}

func (a *articleV2) GenerateBodyByAI(ctx context.Context, prompt string) (string, error) {

	// toolの定義
	tools := chatgptDto.Tools{
		&chatgptDto.ToolFunction{
			Type: "function",
			Function: &chatgptDto.Function{
				Name:        "searchGoogle",
				Description: "When you need information that you don't know or when you need the latest information, search the internet (Google) for it.",
				Parameters: chatgptDto.FuncitonParameters{
					Type: "object",
					Properties: map[string]*chatgptDto.FunctionPropertiesValue{
						"keyword": {
							Type: "string",
						},
					},
					Required: []string{"keyword"},
				},
			},
		},
	}

	// 一度目
	output, err := a.chatGPTClient.CreateChatCompletions(ctx, &chatgptDto.ChatCompletionsInput{
		Model: chatgpt.GPT4Model,
		Messages: []chatgptDto.Message{
			&chatgptDto.SystemMessage{
				Role:    "system",
				Content: "You are a highly knowledgeable assistant.",
			},
			&chatgptDto.SystemMessage{
				Role:    "system",
				Content: "Write detailed articles in Markdown format.",
			},
			&chatgptDto.SystemMessage{
				Role:    "system",
				Content: "Avoid using code blocks in your Markdown responses.",
			},
			&chatgptDto.SystemMessage{
				Role:    "user",
				Content: prompt,
			},
		},
		Tools:      tools,
		ToolChoice: "auto",
	})
	if err != nil {
		return "", fmt.Errorf("failed chatGPTClient.CreateChatCompletions. err: %w", err)
	}

	if len(output.Choices) == 0 {
		return "", fmt.Errorf("empty result")
	}

	// tool callsではない場合はそのまま返す
	if output.Choices[0].FinishReason != "tool_calls" {
		return output.GetMessage(), nil
	}

	// google検索
	// TODO これだと不十分だから、サマライズのAPIを利用して、中身を取るのがいいかも
	arguments := output.Choices[0].Message.ToolCalls[0].Function.Arguments
	argmentsMap := map[string]string{}
	if err := json.Unmarshal([]byte(arguments), &argmentsMap); err != nil {
		return "", fmt.Errorf("failed json.Unmarshal. err: %w", err)
	}
	googleSearchResults, err := a.googleCustomSearchRepo.Search(ctx, argmentsMap["keyword"])
	if err != nil {
		return "", fmt.Errorf("failed google search. err: %w", err)
	}

	googleSearchResultsJSON, err := json.Marshal(googleSearchResults)
	if err != nil {
		return "", fmt.Errorf("failed json.Unmarshal. err: %w", err)
	}

	toolCallID := output.Choices[0].Message.ToolCalls[0].ID

	// まとめ
	output2, err := a.chatGPTClient.CreateChatCompletions(ctx, &chatgptDto.ChatCompletionsInput{
		Model: chatgpt.GPT4Model,
		Messages: []chatgptDto.Message{
			&chatgptDto.AssistantMessage{
				Role:    "assistant",
				Content: "tool_calls",
				ToolCalls: []chatgptDto.AssistantMessageToolCall{
					{
						ID:   toolCallID,
						Type: "function",
						Function: chatgptDto.AssistantMessageToolCallFunction{
							Name:      "searchGoogle",
							Arguments: arguments,
						},
					},
				},
			},
			&chatgptDto.ToolMessage{
				Role:       "tool",
				Content:    string(googleSearchResultsJSON), // google検索結果
				ToolCallID: toolCallID,
			},
			&chatgptDto.SystemMessage{
				Role:    "system",
				Content: "Please provide the user with the information obtained from a Google search and the reference URLs.",
			},
		},
		Tools:      tools,
		ToolChoice: "none",
	})
	if err != nil {
		return "", fmt.Errorf("failed 2回目 chatGPTClient.CreateChatCompletions. err: %w", err)
	}

	return output2.GetMessage(), nil
}
