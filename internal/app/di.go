package app

import (
	"context"
	"log"
	"os"
	"task-service/internal/api"
	task_service2 "task-service/internal/repository/task_service"
	"task-service/internal/usecase/task_service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type diContainer struct {
	server     *api.TaskServiceServer
	useCase    api.TaskServiceUsecase
	repository task_service.TaskServiceRepository

	pool *pgxpool.Pool
}

func NewdiContainer() *diContainer {
	return &diContainer{}
}
func (d *diContainer) Server(ctx context.Context) *api.TaskServiceServer {
	if d.server == nil {
		d.server = api.NewTaskServiceServer(d.UseCase(ctx))
	}
	return d.server
}

func (d *diContainer) UseCase(ctx context.Context) api.TaskServiceUsecase {
	if d.useCase == nil {
		d.useCase = task_service.NewTaskServiceUseCase(d.Repository(ctx))
	}
	return d.useCase
}
func (d *diContainer) Repository(ctx context.Context) task_service.TaskServiceRepository {
	if d.repository == nil {
		d.repository = task_service2.NewTaskServiceRepository(d.Pool(ctx))
	}
	return d.repository
}
func (d *diContainer) Pool(ctx context.Context) *pgxpool.Pool {
	if d.pool == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		pool, err := pgxpool.New(ctx, os.Getenv("DB_URI"))
		if err != nil {
			log.Fatalf("Failed to close connection with db: %v", err)
		}
		if err = pool.Ping(ctx); err != nil {
			log.Fatalf("Failed to ping db: %v", err)
		}
		return pool
	}

	return d.pool
}
