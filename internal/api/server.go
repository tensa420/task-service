package api

import (
	"task-service/internal/usecase"
	"task-service/pkg/task_service"
)

type TaskServiceServer struct {
	useCase usecase.TaskServiceUsecase
	task_service.UnimplementedTaskServiceServer
}

func NewTaskServiceServer(useCase usecase.TaskServiceUsecase) *TaskServiceServer {
	return &TaskServiceServer{useCase: useCase}
}
