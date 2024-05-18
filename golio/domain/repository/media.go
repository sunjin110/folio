package repository

import (
	"context"
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Media interface {
	Insert(ctx context.Context, txTime time.Time, id string, fileType string) (uploadPresignedURL string, err error)
	Get(ctx context.Context, id string) (*model.Medium, error)
	Delete(ctx context.Context, id string) error
	FindSummary(ctx context.Context, paging *Paging) ([]*model.MediumSummary, error)
	TotalCount(ctx context.Context) (int32, error)
}
