package dto

// https://platform.openai.com/docs/api-reference/chat/create
type ChatCompletionsInput struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}