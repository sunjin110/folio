package repository

import "context"

type LineMessage interface {
	ReplyMessage(ctx context.Context, replyToken string, message string) error
}
