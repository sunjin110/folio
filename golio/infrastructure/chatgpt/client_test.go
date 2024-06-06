package chatgpt_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/infrastructure/chatgpt"
	"github.com/sunjin110/folio/golio/infrastructure/chatgpt/dto"
)

// go test -v -count=1 -timeout 30s -run ^TestClient_CreateChatCompletions_Real$ github.com/sunjin110/folio/golio/infrastructure/chatgpt
func TestClient_CreateChatCompletions_Real(t *testing.T) {
	SkipConvey("TestClient_CreateChatCompletions_Real", t, func() {

		client := chatgpt.NewClient("dummy")
		output, err := client.CreateChatCompletions(context.Background(), &dto.ChatCompletionsInput{
			Model: "gpt-3.5-turbo",
			Messages: []dto.Message{
				&dto.SystemMessage{
					Role:    "system",
					Content: "You are a helpful assistant",
				},
				&dto.UserMessage{
					Role:    "user",
					Content: "hello world",
				},
			},
		})
		So(err, ShouldBeNil)

		b, err := json.Marshal(output)
		So(err, ShouldBeNil)
		fmt.Println("output is ", string(b))

	})
}
