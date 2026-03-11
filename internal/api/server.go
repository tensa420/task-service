package api

import (
	"task-service/pkg/task_service"
)

type TaskServiceServer struct {
	useCase TaskServiceUsecase
	task_service.UnimplementedTaskServiceServer
}

func NewTaskServiceServer(useCase TaskServiceUsecase) *TaskServiceServer {
	return &TaskServiceServer{useCase: useCase}
}
