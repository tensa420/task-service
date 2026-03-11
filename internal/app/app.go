package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"task-service/pkg/task_service"

	"google.golang.org/grpc"
)

type App struct {
	diContainer *diContainer
	grpcServer  *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	err := a.initDI(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewdiContainer()
	return nil
}
func (a *App) RunGRPCServer(ctx context.Context) error {
	grpcServer := a.diContainer.Server(ctx)
	srv := grpc.NewServer()
	task_service.RegisterTaskServiceServer(srv, grpcServer)

	lis, err := net.Listen("tcp", net.JoinHostPort(os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT")))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Printf("server started on %s", lis.Addr().String())
	if err = srv.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func (a *App) Run(ctx context.Context) error {
	err := a.RunGRPCServer(ctx)
	if err != nil {
		return err
	}
	return nil
}
