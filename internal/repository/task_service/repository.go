package task_service

import (
	"context"
	"task-service/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskServiceRepo struct {
	pool pgxpool.Pool
}

func NewTaskServiceRepository(pool pgxpool.Pool) *TaskServiceRepo {
	return &TaskServiceRepo{pool: pool}
}

func (r *TaskServiceRepo) CreateTask(ctx context.Context, task entity.Task) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO tasks(task_uuid,user_uuid,title,description,status) VALUES $1, $2, $3,$4,$5", task.TaskUUID, task.UserUUID, task.Title, task.Description, task.Status)
	if err != nil {
		return err
	}
	return nil
}
func (r *TaskServiceRepo) GetTask(ctx context.Context, taskUUID, userUUID string) (entity.Task, error) {
	var task entity.Task
	row := r.pool.QueryRow(ctx, "SELECT * FROM tasks WHERE task_uuid = $1 and user_uuid = $2", taskUUID, userUUID)
	err := row.Scan(
		&task.TaskUUID,
		&task.UserUUID,
		&task.Title,
		&task.Description,
		&task.Status)
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}
func (r *TaskServiceRepo) FinishTask(ctx context.Context, taskUUID, userUUID string) error {
	_, err := r.pool.Exec(ctx, "UPDATE tasks SET status = $1 WHERE task_uuid = $2 and user_uuid = $3",
		entity.TaskStatusFinished, taskUUID, userUUID)
	if err != nil {
		return err
	}
	return nil
}
func (r *TaskServiceRepo) DeleteTask(ctx context.Context, taskUUID, userUUID string) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM tasks WHERE task_uuid = $1 and user_uuid = $2", taskUUID, userUUID)
	if err != nil {
		return err
	}
	return nil
}
func (r *TaskServiceRepo) GetListOfTasks(ctx context.Context, userUUID string) ([]entity.Task, error) {
	var tasks []entity.Task
	rows, err := r.pool.Query(ctx, "SELECT * FROM tasks WHERE user_uuid = $1", userUUID)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		var task entity.Task
		rows.Scan(
			&task.TaskUUID,
			&task.UserUUID,
			&task.Title,
			&task.Description,
			&task.Status)
		tasks = append(tasks, task)
	}
	return tasks, nil
}
