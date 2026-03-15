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

func (u *TaskUseCase) DeleteTask(ctx context.Context, taskUUID string, userUUID string) error {
	taskLog := events.TaskLog{
		LogUUID:    uuid.New(),
		UserUUID:   uuid.MustParse(userUUID),
		TaskUUID:   uuid.MustParse(taskUUID),
		Created_at: time.Now(),
		Type:       events.LogTypeDelete,
	}
	err := u.repo.SelectForUpdate(ctx, taskUUID, userUUID, func(tx *gorm.DB, task entity.Task) error {
		if err := tx.Delete(&task).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return entity.ErrNotFound
			}
		}
		if err := u.logproducer.SendMessage(ctx, taskLog); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
