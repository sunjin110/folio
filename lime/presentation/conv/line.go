package conv

import (
	"log/slog"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/sunjin110/folio/lime/domain/model"
)

func ToLineEventsModel(events []webhook.EventInterface) model.LineEvents {
	models := make(model.LineEvents, 0, len(events))
	for _, event := range events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			models = append(models, toLineMessageEvent(&e))
		case *webhook.MessageEvent:
			models = append(models, toLineMessageEvent(e))
		}
	}
	return models
}

func toLineMessageEvent(event *webhook.MessageEvent) *model.LineMessageEvent {
	if event == nil {
		return nil
	}
	return &model.LineMessageEvent{
		LineMessageContent: toLineMessageContent(event.Message),
	}
}

func toLineMessageContent(messageContent webhook.MessageContentInterface) *model.LineMessageContent {
	switch c := messageContent.(type) {
	case webhook.ImageMessageContent:
		return &model.LineMessageContent{
			ID:                  c.Id,
			LineContentProvider: toLineContentProvider(c.Id, c.ContentProvider),
		}
	case webhook.AudioMessageContent:
		return &model.LineMessageContent{
			ID:                  c.Id,
			LineContentProvider: toLineContentProvider(c.Id, c.ContentProvider),
		}
	case webhook.VideoMessageContent:
		return &model.LineMessageContent{
			ID:                  c.Id,
			LineContentProvider: toLineContentProvider(c.Id, c.ContentProvider),
		}
	default:
		slog.Info("対応していないcontentです", "messageContent", messageContent)
	}
	return nil
}

func toLineContentProvider(messageID string, provider *webhook.ContentProvider) model.LineContentProvider {
	if provider == nil {
		return nil
	}

	switch provider.Type {
	case webhook.ContentProviderTYPE_EXTERNAL:
		return &model.LineContentProviderExternal{
			OriginalContentURL: provider.OriginalContentUrl,
			PreviewImageURL:    provider.PreviewImageUrl,
		}
	case webhook.ContentProviderTYPE_LINE:
		return &model.LineContentProviderLine{
			MessageID: messageID,
		}
	default:
		return nil
	}
}
