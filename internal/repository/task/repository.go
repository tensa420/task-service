package task

import (
	"context"
	"task-service/internal/entity"

	"gorm.io/gorm"
)

type TaskServiceRepo struct {
	db *gorm.DB
}

func NewTaskServiceRepository(db *gorm.DB) *TaskServiceRepo {
	return &TaskServiceRepo{db: db}
}

func (r *TaskServiceRepo) CreateTask(ctx context.Context, task entity.Task) error {
	return r.db.WithContext(ctx).Create(&task).Error
}

func (r *TaskServiceRepo) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	var task entity.Task
	result := r.db.WithContext(ctx).
		Where("task_uuid = ? AND user_uuid = ?", taskUUID, userUUID).
		First(&task)

	if result.Error != nil {
		return entity.Task{}, result.Error
	}

	return task, nil
}

func (r *TaskServiceRepo) FinishTask(ctx context.Context, taskUUID, userUUID string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Task{}).
		Where("task_uuid = ? AND user_uuid = ?", taskUUID, userUUID).
		Update("status", entity.TaskStatusFinished).Error
}

func (r *TaskServiceRepo) DeleteTask(ctx context.Context, taskUUID, userUUID string) error {
	return r.db.WithContext(ctx).
		Where("task_uuid = ? AND user_uuid = ?", taskUUID, userUUID).
		Delete(&entity.Task{}).Error
}

func (r *TaskServiceRepo) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	var tasks []entity.Task
	result := r.db.WithContext(ctx).
		Where("user_uuid = ?", userUUID).
		Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
