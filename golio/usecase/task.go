package usecase

import (
	"context"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type Task interface {
	Get(ctx context.Context, id string) (*model.Task, error)
	Insert(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id string) error
}

type task struct {
	taskRepo repository.Task
}

func NewTask(taskRepo repository.Task) Task {
	return &task{
		taskRepo: taskRepo,
	}
}

func (t *task) Get(ctx context.Context, id string) (*model.Task, error) {
	task, err := t.taskRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed taskRepo.Get. err: %w", err)
	}
	return task, nil
}

func (t *task) Delete(ctx context.Context, id string) error {
	if err := t.taskRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed taskRepo.Delete. err: %w", err)
	}
	return nil
}

func (t *task) Insert(ctx context.Context, task *model.Task) error {
	if err := t.taskRepo.Insert(ctx, task); err != nil {
		return fmt.Errorf("failed taskRepo.Insert. err: %w", err)
	}
	return nil
}

func (t *task) Update(ctx context.Context, task *model.Task) error {
	if err := t.taskRepo.Update(ctx, task); err != nil {
		return fmt.Errorf("failed taskRepo.Update. err: %w", err)
	}
	return nil
}
