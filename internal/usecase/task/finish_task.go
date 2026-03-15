package task

import (
	"context"
	"task-service/internal/entity"
	"task-service/internal/entity/events"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *TaskUseCase) FinishTask(ctx context.Context, taskUUID, userUUID string) error {
	taskLog := events.TaskLog{
		TaskUUID:   uuid.MustParse(taskUUID),
		UserUUID:   uuid.MustParse(userUUID),
		Created_at: time.Now(),
		LogUUID:    uuid.New(),
		Type:       events.LogTypeFinish,
	}
	err := u.repo.SelectForUpdate(ctx, taskUUID, userUUID, func(tx *gorm.DB, task entity.Task) error {
		if err := tx.Model(&task).Update("status", entity.TaskStatusFinished).Error; err != nil {
			return err
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
