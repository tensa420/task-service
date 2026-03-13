package task

import (
	"context"
	"errors"
	"task-service/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (u *TaskServiceUseCase) DeleteTask(ctx context.Context, taskUUID string, userUUID string) error {
	err := u.repo.DeleteTask(ctx, taskUUID, userUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.ErrNotFound
		}
		return err
	}
	return nil
}
