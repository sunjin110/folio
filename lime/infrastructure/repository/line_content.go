package repository

import (
	"context"
	"fmt"
	"mime"
	"net/http"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/sunjin110/folio/lime/domain/model"
	"github.com/sunjin110/folio/lime/domain/repository"
)

type lineContent struct {
	lineMessageClient *messaging_api.MessagingApiAPI
	lineClient        *linebot.Client
}

func NewLineContent(lineMessageClient *messaging_api.MessagingApiAPI) repository.LineContent {
	return &lineContent{
		lineMessageClient: lineMessageClient,
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
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("failed mime.ParseMediaType. err: %w", err)
	}

	return model.NewContent(resp.Body, mediatype, resp.ContentLength), nil
}

func (l *lineContent) getFromLine(ctx context.Context, provider *model.LineContentProviderLine) (*model.Content, error) {
	resp, err := l.lineClient.GetMessageContent(provider.MessageID).WithContext(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed lineClient.GetMessageContent. messageID: %s, err: %w", provider.MessageID, err)
	}

	mediatype, _, err := mime.ParseMediaType(resp.ContentType)
	if err != nil {
		return nil, fmt.Errorf("failed mime.ParseMediaType. err: %w", err)
	}
	return model.NewContent(resp.Content, mediatype, resp.ContentLength), nil
}
