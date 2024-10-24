package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/sunjin110/folio/lime/domain/model"
	"github.com/sunjin110/folio/lime/domain/repository"
)

type lineContent struct {
	lineClient *linebot.Client
}

func NewLineContent(lineClient *linebot.Client) repository.LineContent {
	return &lineContent{
		lineClient: lineClient,
	}
}

func (l *lineContent) GetByProvider(ctx context.Context, provider model.LineContentProvider) (*model.Content, error) {
	switch p := provider.(type) {
	case *model.LineContentProviderExternal:
		return l.getFromExternal(p)
	case *model.LineContentProviderLine:
		return l.getFromLine(ctx, p)
	}
	return nil, fmt.Errorf("undefined provider type. type: %d", provider.GetType())
}

func (*lineContent) getFromExternal(provider *model.LineContentProviderExternal) (*model.Content, error) {
	resp, err := http.Get(provider.OriginalContentURL)
	if err != nil {
		return nil, fmt.Errorf("failed http.Get. OriginalContentURL: %s, err: %w", provider.OriginalContentURL, err)
	}
	contentType := resp.Header.Get("Content-Type")
	return model.NewContent(resp.Body, contentType, resp.ContentLength, nil), nil
}

func (l *lineContent) getFromLine(ctx context.Context, provider *model.LineContentProviderLine) (*model.Content, error) {
	resp, err := l.lineClient.GetMessageContent(provider.MessageID).WithContext(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed lineClient.GetMessageContent. messageID: %s, err: %w", provider.MessageID, err)
	}

	return model.NewContent(resp.Content, resp.ContentType, resp.ContentLength, nil), nil
}
