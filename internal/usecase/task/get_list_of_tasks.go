package task

import (
	"context"
	"errors"
	"fmt"
	"os"
	"task-service/internal/entity"

	"gorm.io/gorm"
)

func (u *TaskServiceUseCase) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	tasks, err := u.repo.GetListOfTasks(ctx, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Task{}, entity.ErrNotFound
		}
		return []entity.Task{}, entity.ErrNotFound
	}

	producerMsg := fmt.Sprintf("User %v got his list of tasks %v", userUUID, tasks)
	err = u.logproducer.SendMessage(ctx, os.Getenv("KAFKA_TOPIC_LOGS"), userUUID, producerMsg)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
