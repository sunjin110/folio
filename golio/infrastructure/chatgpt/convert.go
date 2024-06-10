package chatgpt

import (
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func GetMessageContent(resp openai.ChatCompletionResponse) string {
	contents := []string{}
	for _, choice := range resp.Choices {
		contents = append(contents, choice.Message.Content)
	}
	return strings.Join(contents, "\n")
}
