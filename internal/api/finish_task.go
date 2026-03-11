package api

import (
	"context"
	"task-service/pkg/task_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *TaskServiceServer) FinishTask(ctx context.Context, in *task_service.FinishTaskRequest) (*emptypb.Empty, error) {
	err := s.useCase.FinishTask(ctx, in.TaskUUID, in.UserUUID)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
