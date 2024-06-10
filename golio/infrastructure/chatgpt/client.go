package chatgpt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type Client interface {
	CreateChatCompletions(ctx context.Context, input openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error)
}

type client struct {
	client *openai.Client
}

func NewClient(apiKey string) Client {
	return &client{
		client: openai.NewClient(apiKey),
	}
}

func (c *client) CreateChatCompletions(ctx context.Context, input openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	output, err := c.client.CreateChatCompletion(ctx, input)
	if err != nil {
		return openai.ChatCompletionResponse{}, fmt.Errorf("failed client.CreateChatCompletion")
	}
	return output, nil
}
