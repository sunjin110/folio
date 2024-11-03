package application

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/sunjin110/folio/gomb/domain/repository"
)

const thumbnailPathPrefix = "thumbnail"

type Thumbnail interface {
	GenerateAndPutThumbnail(ctx context.Context, s3Key string) error
}

type thumbnailApplication struct {
	storageRepo repository.Storage
}

func NewThumbanilApplication(storageRepo repository.Storage) Thumbnail {
	return &thumbnailApplication{
		storageRepo: storageRepo,
	}
}

func (t *thumbnailApplication) GenerateAndPutThumbnail(ctx context.Context, s3Key string) error {

	if strings.HasPrefix(s3Key, thumbnailPathPrefix) {
		slog.Info("this is a thumbnail", "s3Key", s3Key)
		return nil
	}

	originalContent, err := t.storageRepo.GetContent(ctx, s3Key)
	if err != nil {
		return fmt.Errorf("failed get original content. err: %w", err)
	}

	if !originalContent.IsImage() && !originalContent.IsVideo() {
		slog.Info("this is not media. skip making thumbnail", "fileName", originalContent.FileName())
		return nil
	}

	thumbnail, err := originalContent.Thumbnail()
	if err != nil {
		return fmt.Errorf("failed generate thumbnail. s3Key: %s, err: %w", s3Key, err)
	}

	thumbnailContent, err := thumbnail.ToContent(s3Key)
	if err != nil {
		return fmt.Errorf("failed thumbnail.ToContent. s3Key: %s, err: %w", s3Key, err)
	}

	if err := t.storageRepo.SaveContent(ctx, thumbnailPathPrefix, thumbnailContent); err != nil {
		return fmt.Errorf("failed save thumbnail. thumbnailPathPrefix: %s, err: %w", thumbnailPathPrefix, err)
	}

	return nil
}
