package usecase

import (
	"context"
	"task-service/internal/entity"
)

type TaskServiceTypes interface {
	CreateTask(ctx context.Context, task entity.Task) (string, error)
	DeleteTask(ctx context.Context, ID, userID string) error
	FinishTask(ctx context.Context, taskID, userID string) error
	GetListOfTasks(ctx context.Context, userID string) ([]entity.Task, error)
	GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error)
}
