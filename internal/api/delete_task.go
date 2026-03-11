package api

import (
	"context"
	"task-service/pkg/task_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *TaskServiceServer) DeleteTask(ctx context.Context, in *task_service.DeleteTaskRequest) (*emptypb.Empty, error) {
	err := s.useCase.DeleteTask(ctx, in.TaskUUID, in.UserUUID)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
