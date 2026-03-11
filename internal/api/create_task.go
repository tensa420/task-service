package api

import (
	"context"
	"task-service/internal/entity"
	"task-service/pkg/task_service"

	"github.com/google/uuid"
)

func (s *TaskServiceServer) CreateTask(ctx context.Context, in *task_service.CreateTaskRequest) (*task_service.CreateTaskResponse, error) {
	taskUUID, err := s.useCase.CreateTask(ctx, convertProtoToEntity(in.Task))
	if err != nil {
		return nil, err
	}
	return &task_service.CreateTaskResponse{TaskUUID: taskUUID}, nil
}

func convertProtoToEntity(task *task_service.Task) entity.Task {
	return entity.Task{
		UserUUID:    uuid.MustParse(task.UserUUID),
		Description: task.Description,
		Title:       task.Title,
	}
}
