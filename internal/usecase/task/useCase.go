package task

import "task-service/internal/producer"

type TaskUseCase struct {
	repo        TaskRepository
	logproducer producer.LogProducer
}

func NewTaskServiceUseCase(repo TaskRepository, logProducer producer.LogProducer) *TaskUseCase {
	return &TaskUseCase{
		repo:        repo,
		logproducer: logProducer,
	}
}
