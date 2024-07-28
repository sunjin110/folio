package conv

import (
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
)

func NewTasks(tasks []*postgres_dto.Task, detials []*postgres_dto.TaskDetail) model.Tasks {
	detailMap := make(map[string]*postgres_dto.TaskDetail, len(detials))
	for _, detail := range detials {
		detailMap[detail.ID] = detail
	}

	models := make(model.Tasks, 0, len(tasks))
	for _, task := range tasks {
		detail, ok := detailMap[task.ID]
		if !ok {
			continue
		}
		models = append(models, NewTask(task, detail))
	}
	return models
}

func NewTask(task *postgres_dto.Task, detail *postgres_dto.TaskDetail) *model.Task {
	return &model.Task{
		ID:        task.ID,
		Title:     task.Title,
		Detail:    detail.Detail,
		Status:    model.TaskStatus(task.Status),
		StartTime: task.StartTime,
		DueTime:   task.DueTime,
		CratedAt:  task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}
