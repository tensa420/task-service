package task_service

import (
	"context"
	"errors"
	"task-service/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (u *TaskServiceUseCase) FinishTask(ctx context.Context, taskUUID, userUUID string) error {
	err := u.repo.FinishTask(ctx, taskUUID, userUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.ErrNotFound
		}
		return err
	}
	return nil
}
