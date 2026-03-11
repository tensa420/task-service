package api

import (
	"context"
	"task-service/internal/entity"
)

type TaskServiceUsecase interface {
	CreateTask(ctx context.Context, task entity.Task) (string, error)
	DeleteTask(ctx context.Context, TaskUUID, userUUID string) error
	FinishTask(ctx context.Context, taskUUID, userUUID string) error
	GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error)
	GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error)
}
