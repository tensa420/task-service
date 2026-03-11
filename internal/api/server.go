package api

import (
	"task-service/internal/usecase"
	"task-service/pkg/task_service"
)

type TaskServiceServer struct {
	useCase usecase.TaskServiceTypes
	task_service.UnimplementedTaskServiceServer
}

func NewTaskServiceServer(useCase usecase.TaskServiceTypes) *TaskServiceServer {
	return &TaskServiceServer{useCase: useCase}
}
