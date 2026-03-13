package task

import (
	"context"
	"errors"
	"task-service/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (u *TaskServiceUseCase) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	tasks, err := u.repo.GetListOfTasks(ctx, userUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}
	return tasks, nil
}
