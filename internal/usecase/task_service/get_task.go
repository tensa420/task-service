package task_service

import (
	"context"
	"errors"
	"task-service/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (u *TaskServiceUseCase) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	task, err := u.repo.GetTask(ctx, taskUUID, userUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Task{}, entity.ErrNotFound
		}
		return entity.Task{}, err
	}
	return task, nil
}
