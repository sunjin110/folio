package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TaskStatus int

const (
	TaskStatusTodo TaskStatus = 1
	TaskStatusDone TaskStatus = 2
)

type Tasks []*Task

// Task 仕事のタスクです
type Task struct {
	ID        string
	Title     string
	Detail    string
	Status    TaskStatus
	StartTime *time.Time
	DueTime   *time.Time
	CratedAt  time.Time
	UpdatedAt time.Time
}

func NewTask(title string, detail string, startTime *time.Time, dueTime *time.Time, now time.Time) (*Task, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed uuid.NewRandom. err: %w", err)
	}

	return &Task{
		ID:        u.String(),
		Title:     title,
		Detail:    detail,
		Status:    TaskStatusTodo,
		StartTime: startTime,
		DueTime:   dueTime,
		CratedAt:  now,
		UpdatedAt: now,
	}, nil
}
