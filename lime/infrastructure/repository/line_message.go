package repository

import (
	"context"
	"fmt"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/sunjin110/folio/lime/domain/repository"
)

type lineMessage struct {
	lineMessageClient *messaging_api.MessagingApiAPI
}

func NewLineMessage(lineMessageClient *messaging_api.MessagingApiAPI) repository.LineMessage {
	return &lineMessage{
		lineMessageClient: lineMessageClient,
	}
}

func (l *lineMessage) ReplyMessage(ctx context.Context, replyToken string, message string) error {
	_, err := l.lineMessageClient.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: replyToken,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: message,
			},
		},
		NotificationDisabled: false,
	})
	if err != nil {
		return fmt.Errorf("failed lineMessageClient.ReplyMessage. err: %w", err)
	}
	return nil
}
