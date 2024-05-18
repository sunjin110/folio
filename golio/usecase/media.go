package usecase

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type Media interface {
	Get(ctx context.Context, id string) (*model.Medium, error)
	Insert(ctx context.Context, fileName string) (updateURL string, err error)
	Delete(ctx context.Context, id string) error
	FindSummaries(ctx context.Context, offset int32, limit int32) (*FindMediaSummariesOutput, error)
}

type FindMediaSummariesOutput struct {
	Media      []*model.MediumSummary
	TotalCount int32
}

type media struct {
	mediaRepo repository.Media
}

func NewMedia(mediaRepo repository.Media) Media {
	return &media{
		mediaRepo: mediaRepo,
	}
}

func (m *media) Delete(ctx context.Context, id string) error {
	if err := m.mediaRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed medium delete. id: %s, err: %w", id, err)
	}
	return nil
}

func (m *media) FindSummaries(ctx context.Context, offset int32, limit int32) (*FindMediaSummariesOutput, error) {
	summaries, err := m.mediaRepo.FindSummary(ctx, &repository.Paging{
		Offset: int(offset),
		Limit:  int(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("failed find media summaries: %w", err)
	}

	totalCount, err := m.mediaRepo.TotalCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed total media count: %w", err)
	}
	return &FindMediaSummariesOutput{
		Media:      summaries,
		TotalCount: totalCount,
	}, nil
}

func (m *media) Get(ctx context.Context, id string) (*model.Medium, error) {
	medium, err := m.mediaRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed get medium. id: %s, err: %w", id, err)
	}
	return medium, nil
}

func (m *media) Insert(ctx context.Context, filteName string) (string, error) {
	txTime := time.Now()

	fileType := filepath.Ext(filteName)
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed new uuid: %w", err)
	}
	uploadPresignedURL, err := m.mediaRepo.Insert(ctx, txTime, id.String(), fileType)
	if err != nil {
		return "", fmt.Errorf("failed media insert: %w", err)
	}

	return uploadPresignedURL, nil
}
