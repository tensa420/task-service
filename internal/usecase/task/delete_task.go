package task

import (
	"context"
	"errors"
	"fmt"
	"os"
	"task-service/internal/entity"

	"gorm.io/gorm"
)

func (u *TaskServiceUseCase) DeleteTask(ctx context.Context, taskUUID string, userUUID string) error {
	err := u.repo.DeleteTask(ctx, taskUUID, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ErrNotFound
		}
		return err
	}
	producerMsg := fmt.Sprintf("User %v deleted task %v", userUUID, taskUUID)

	err = u.logproducer.SendMessage(ctx, os.Getenv("KAFKA_TOPIC_LOGS"), userUUID, producerMsg)
	if err != nil {
		return err
	}

	return nil
}
