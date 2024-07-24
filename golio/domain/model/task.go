package model

import "time"

type TaskStatus int

const (
	TaskStatusTodo TaskStatus = 1
	TaskStatusDone TaskStatus = 2
)

type Tasks []*Task

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
