package task

import (
	"context"
	"errors"
	"task-service/internal/entity"
	"task-service/internal/entity/events"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *TaskUseCase) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	task.TaskUUID = uuid.New()
	task.Status = entity.TaskStatusNew

	err := u.repo.CreateTask(ctx, task)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", entity.ErrNotFound
		}
	}

	taskLog := events.TaskLog{
		TaskUUID:   task.TaskUUID,
		UserUUID:   task.UserUUID,
		Created_at: time.Now(),
		LogUUID:    uuid.New(),
		Type:       events.LogTypeCreate,
	}
	
	err = u.logproducer.SendMessage(ctx, taskLog)
	if err != nil {
		return "", err
	}

	return task.TaskUUID.String(), nil
}
