package application

import (
	"context"
	"fmt"

	"github.com/sunjin110/folio/lime/domain/model"
	"github.com/sunjin110/folio/lime/domain/repository"
)

type LineUsecase interface {
	SaveContents(ctx context.Context, events model.LineEvents) error
}

type lineUsecase struct {
	lineContentRepo repository.LineContent
	lineMessageRepo repository.LineMessage
	storageRepo     repository.Storage
}

func NewLineUsecase(lineContentRepo repository.LineContent, lineMessageRepo repository.LineMessage, storageRepo repository.Storage) LineUsecase {
	return &lineUsecase{
		lineContentRepo: lineContentRepo,
		lineMessageRepo: lineMessageRepo,
		storageRepo:     storageRepo,
	}
}

func (l *lineUsecase) SaveContents(ctx context.Context, events model.LineEvents) error {
	for _, event := range events {
		switch e := event.(type) {
		case *model.LineMessageEvent:
			content, err := l.lineContentRepo.GetByProvider(ctx, e.LineMessageContent.LineContentProvider)
			if err != nil {
				return fmt.Errorf("failed lineContentRepo.GetByProvider. err: %w", err)
			}
			if err := l.storageRepo.SaveContent(ctx, content); err != nil {
				return fmt.Errorf("failed storageRepo.SaveContent. err: %w", err)
			}
			if err := l.lineMessageRepo.ReplyMessage(ctx, e.ReplyToken, "success upload"); err != nil {
				return fmt.Errorf("failed lineMessageRepo.ReplyMessage. err: %w", err)
			}
		}
	}
	return nil
}
