package task

import "task-service/internal/producer"

type TaskServiceUseCase struct {
	repo        TaskServiceRepository
	logproducer producer.LogProducer
}

func NewTaskServiceUseCase(repo TaskServiceRepository) *TaskServiceUseCase {
	return &TaskServiceUseCase{repo: repo}
}
