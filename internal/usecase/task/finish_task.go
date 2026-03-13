package task

import (
	"context"
	"errors"
	"fmt"
	"os"
	"task-service/internal/entity"

	"gorm.io/gorm"
)

func (u *TaskServiceUseCase) FinishTask(ctx context.Context, taskUUID, userUUID string) error {
	err := u.repo.FinishTask(ctx, taskUUID, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ErrNotFound
		}
		return err
	}

	producerMsg := fmt.Sprintf("User %v finished task %v", userUUID, taskUUID)
	err = u.logproducer.SendMessage(ctx, os.Getenv("KAFKA_TOPIC_LOGS"), userUUID, producerMsg)
	if err != nil {
		return err
	}

	return nil
}
