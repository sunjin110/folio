package postgres_dto

import "time"

type Task struct {
	ID        string     `db:"id"`
	Title     string     `db:"title"`
	Status    int        `db:"status"`
	StartTime *time.Time `db:"start_time"`
	DueTime   *time.Time `db:"due_time"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

type TaskDetail struct {
	ID        string    `db:"id"`
	TaskID    string    `db:"task_id"`
	Detail    string    `db:"detail"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
