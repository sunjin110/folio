package dto

import "strings"

type ChatCompletionsOutput struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Model   string    `json:"model"`
	Choices []*Choise `json:"choices"`
}

func (output *ChatCompletionsOutput) GetMessage() string {
	contents := []string{}
	for _, choice := range output.Choices {
		if choice == nil || choice.Message == nil || choice.Message.Content == nil {
			continue
		}
		contents = append(contents, *choice.Message.Content)
	}
	return strings.Join(contents, "\n")
}
