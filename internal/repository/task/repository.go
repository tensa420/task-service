package task

import (
	"context"
	"task-service/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskServiceRepository(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(ctx context.Context, task entity.Task) error {
	return r.db.WithContext(ctx).Create(&task).Error
}

func (r *TaskRepo) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	var task entity.Task
	result := r.db.WithContext(ctx).
		Where("task_uuid = ? AND user_uuid = ?", taskUUID, userUUID).
		First(&task)

	if result.Error != nil {
		return entity.Task{}, result.Error
	}

	return task, nil
}

func (r *TaskRepo) FinishTask(ctx context.Context, taskUUID, userUUID string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Task{}).
		Where("task_uuid = ? AND user_uuid = ?", taskUUID, userUUID).
		Update("status", entity.TaskStatusFinished).Error
}

func (r *TaskRepo) DeleteTask(ctx context.Context, taskUUID, userUUID string) error {
	return r.db.WithContext(ctx).
		Where("task_uuid = ? AND user_uuid = ?", taskUUID, userUUID).
		Delete(&entity.Task{}).Error
}

func (r *TaskRepo) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	var tasks []entity.Task
	result := r.db.WithContext(ctx).
		Where("user_uuid = ?", userUUID).
		Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *TaskRepo) SelectForUpdate(ctx context.Context, taskUUID string, fn func(tx *gorm.DB, task entity.Task) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var task entity.Task
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("task_uuid = ?", taskUUID).
			First(&task).Error; err != nil {
			return err
		}
		return fn(tx, task)
	})
}
