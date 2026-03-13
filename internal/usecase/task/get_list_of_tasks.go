package task

import (
	"context"
	"errors"
	"task-service/internal/entity"

	"gorm.io/gorm"
)

func (u *TaskUseCase) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	tasks, err := u.repo.GetListOfTasks(ctx, userUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Task{}, entity.ErrNotFound
		}
		return []entity.Task{}, entity.ErrNotFound
	}
	return tasks, nil
}
