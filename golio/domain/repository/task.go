package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Task interface {
	Get(ctx context.Context, id string) (*model.Task, error)
	Insert(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, task *model.Task) error
	GetTasks(ctx context.Context) (model.Tasks, error)
	Delete(ctx context.Context, id string) error
}
