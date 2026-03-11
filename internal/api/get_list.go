package api

import (
	"context"
	"task-service/internal/entity"
	"task-service/pkg/task_service"
)

func (s *TaskServiceServer) GetListOfTasks(ctx context.Context, in *task_service.GetListOfTasksRequest) (*task_service.GetListOfTasksResponse, error) {
	tasks, err := s.useCase.GetListOfTasks(ctx, in.GetUserUUID())
	if err != nil {
		return nil, err
	}
	return &task_service.GetListOfTasksResponse{Tasks: convertEntitySliceOfTasksToProto(tasks)}, nil
}

func convertEntitySliceOfTasksToProto(tasks []entity.Task) []*task_service.Task {
	protoTasks := make([]*task_service.Task, 0, len(tasks))
	for _, task := range tasks {
		protoTask := &task_service.Task{
			Title:       task.Title,
			Description: task.Description,
			Status:      convertTaskStatus(task.Status),
			UserUUID:    task.UserUUID.String(),
			TaskUUID:    task.UserUUID.String(),
		}
		protoTasks = append(protoTasks, protoTask)
	}
	return protoTasks
}

func convertTaskStatus(status entity.TaskStatus) task_service.TaskStatus {
	switch status {
	case entity.TaskStatusNew:
		return task_service.TaskStatus_TASK_STATUS_NEW
	default:
		return task_service.TaskStatus_TASK_STATUS_FINISHED
	}
}
