package task_service

import (
	"context"
	"task-service/internal/entity"
)

type TaskServiceRepository interface {
	CreateTask(ctx context.Context, task entity.Task) error
	DeleteTask(ctx context.Context, TaskUUID, userUUID string) error
	FinishTask(ctx context.Context, taskUUID, userUUID string) error
	GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error)
	GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error)
}
