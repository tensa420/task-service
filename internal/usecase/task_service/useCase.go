package task_service

import "task-service/internal/repository"

type TaskServiceUseCase struct {
	repo repository.TaskServiceRepository
}

func NewTaskServiceUseCase(repo repository.TaskServiceRepository) *TaskServiceUseCase {
	return &TaskServiceUseCase{repo: repo}
}
