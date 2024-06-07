package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sunjin110/folio/golio/infrastructure/chatgpt/dto"
)

const endpoint = "https://api.openai.com"

const (
	GPT3Point5TurboModel = "gpt-3.5-turbo"
	GPT4Model            = "gpt-4"
)

const (
	createChatCompletionsPath = "/v1/chat/completions"
)

type Client interface {
	CreateChatCompletions(ctx context.Context, input *dto.ChatCompletionsInput) (*dto.ChatCompletionsOutput, error)
}

type client struct {
	apiKey string
	client *http.Client
}

func NewClient(apiKey string) Client {
	return &client{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (c *client) CreateChatCompletions(ctx context.Context, input *dto.ChatCompletionsInput) (*dto.ChatCompletionsOutput, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed json.Marshal. err: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint+createChatCompletionsPath, strings.NewReader(string(b)))
	if err != nil {
		return nil, fmt.Errorf("failed http.NewRequest. err: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed client.Do. err: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll. err: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed request. status: %d, body: %s, err: %w", resp.StatusCode, string(body), err)
	}

	output := &dto.ChatCompletionsOutput{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal. body: %s, err: %w", string(body), err)
	}
	return output, nil
}
