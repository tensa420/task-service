package api

import (
	"context"
	"task-service/internal/entity"
	"task-service/pkg/task_service"
)

func (s *TaskServiceServer) GetTask(ctx context.Context, in *task_service.GetTaskRequest) (*task_service.GetTaskResponse, error) {
	task, err := s.useCase.GetTask(ctx, in.TaskUUID, in.UserUUID)
	if err != nil {
		return nil, err
	}
	return &task_service.GetTaskResponse{Task: convertEntityTaskToProtoTask(&task)}, nil
}

func convertEntityTaskToProtoTask(task *entity.Task) *task_service.Task {
	protoTask := &task_service.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      convertTaskStatus(task.Status),
		UserUUID:    task.UserUUID.String(),
		TaskUUID:    task.UserUUID.String(),
	}
	return protoTask
}
