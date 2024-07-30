package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/conv"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
	"github.com/sunjin110/folio/golio/infrastructure/repository/query/postgresql"
	"golang.org/x/sync/errgroup"
)

type task struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) repository.Task {
	return &task{
		db: db,
	}
}

func (t *task) Delete(ctx context.Context, id string) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed delete transaction begin. err: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				slog.ErrorContext(ctx, "failed task delete rallback", "err", err)
			}
		}
	}()

	if _, err := t.db.ExecContext(ctx, "delete from tasks where id = $1", id); err != nil {
		return fmt.Errorf("failed delete task. id: %s, err: %w", id, err)
	}

	if _, err := t.db.ExecContext(ctx, "delete from task_details where id = $1", id); err != nil {
		return fmt.Errorf("failed delete task_detail. id: %s, err: %w", id, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed delete task commit. id: %s, err: %w", id, err)
	}

	return nil
}

func (t *task) Get(ctx context.Context, id string) (*model.Task, error) {
	eg, ctx := errgroup.WithContext(ctx)

	detail := &postgres_dto.TaskDetail{}
	eg.Go(func() error {
		if err := t.db.GetContext(ctx, detail, "select * from task_details where id = $1", id); err != nil {
			return fmt.Errorf("failed get task_detail. err: %w", err)
		}
		return nil
	})

	task := &postgres_dto.Task{}
	eg.Go(func() error {
		if err := t.db.GetContext(ctx, task, "select * from tasks where id = $1", id); err != nil {
			return fmt.Errorf("failed get task. err: %w", err)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, fmt.Errorf("failed get task. task_id: %s, err: %w", id, err)
	}
	return conv.NewTask(task, detail), nil
}

func (t *task) GetTasks(ctx context.Context) (model.Tasks, error) {
	eg, ctx := errgroup.WithContext(ctx)

	var tasks []*postgres_dto.Task
	eg.Go(func() error {
		if err := t.db.SelectContext(ctx, tasks, "select * from tasks;"); err != nil {
			return fmt.Errorf("failed select tasks. err: %w", err)
		}
		return nil
	})

	var details []*postgres_dto.TaskDetail
	eg.Go(func() error {
		if err := t.db.SelectContext(ctx, details, "select * from task_details;"); err != nil {
			return fmt.Errorf("failed select task_details. err: %w", err)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, fmt.Errorf("failed get tasks. err: %w", err)
	}
	return conv.NewTasks(tasks, details), nil
}

func (t *task) Insert(ctx context.Context, task *model.Task) error {
	return t.upsert(ctx, task)
}

func (t *task) Update(ctx context.Context, task *model.Task) error {
	return t.upsert(ctx, task)
}

func (t *task) upsert(ctx context.Context, task *model.Task) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed upsert transaction begin. err: %w", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				slog.ErrorContext(ctx, "fialed task upsert rollback", "err", err)
			}
		}
	}()

	if _, err := tx.NamedExecContext(ctx, postgresql.UpsertTask, &postgres_dto.Task{
		ID:        task.ID,
		Title:     task.Title,
		Status:    int(task.Status),
		StartTime: task.StartTime,
		DueTime:   task.DueTime,
		CreatedAt: task.CratedAt,
		UpdatedAt: task.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed upsert tasks. err: %w", err)
	}

	if _, err := tx.NamedExecContext(ctx, postgresql.UpsertTaskDetail, &postgres_dto.TaskDetail{
		ID:        task.ID,
		TaskID:    task.ID,
		Detail:    task.Detail,
		CreatedAt: task.CratedAt,
		UpdatedAt: task.UpdatedAt,
	}); err != nil {
		return fmt.Errorf("failed upsert task_details. err: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed commit. err: %w", err)
	}
	return nil
}
