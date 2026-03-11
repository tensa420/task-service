package task_service

import (
	"context"
	"task-service/internal/entity"

	"github.com/google/uuid"
)

func (u *TaskServiceUseCase) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	task.TaskUUID = uuid.New()
	task.Status = entity.TaskStatusNew
	err := u.repo.CreateTask(ctx, task)
	if err != nil {
		return "", err
	}
	return task.TaskUUID.String(), nil
}
