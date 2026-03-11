package task_service

type TaskServiceUseCase struct {
	repo TaskServiceRepository
}

func NewTaskServiceUseCase(repo TaskServiceRepository) *TaskServiceUseCase {
	return &TaskServiceUseCase{repo: repo}
}
