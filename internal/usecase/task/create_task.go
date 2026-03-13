package task

import (
	"context"
	"errors"
	"fmt"
	"os"
	"task-service/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *TaskServiceUseCase) CreateTask(ctx context.Context, task entity.Task) (string, error) {
	task.TaskUUID = uuid.New()
	task.Status = entity.TaskStatusNew

	err := u.repo.CreateTask(ctx, task)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", entity.ErrNotFound
		}
	}

	producerMsg := fmt.Sprintf("User %v created task %v", task.UserUUID, task.TaskUUID)
	err = u.logproducer.SendMessage(ctx, os.Getenv("KAFKA_TOPIC_LOGS"), task.UserUUID.String(), producerMsg)
	if err != nil {
		return "", err
	}

	return task.TaskUUID.String(), nil
}
