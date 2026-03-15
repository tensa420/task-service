package task

import (
	"context"
	"errors"
	"task-service/internal/entity"

	"gorm.io/gorm"
)

func (u *TaskUseCase) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	task, err := u.repo.GetTask(ctx, taskUUID, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Task{}, entity.ErrNotFound
		}
		return entity.Task{}, err
	}

	return task, nil
}
