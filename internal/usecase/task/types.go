package task

import (
	"context"
	"task-service/internal/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task entity.Task) error
	DeleteTask(ctx context.Context, TaskUUID, userUUID string) error
	FinishTask(ctx context.Context, taskUUID, userUUID string) error
	GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error)
	GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error)
	SelectForUpdate(ctx context.Context, taskUUID string, userUUID string, fn func(tx *gorm.DB, task entity.Task) error) error
}
