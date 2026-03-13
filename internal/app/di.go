package app

import (
	"context"
	"log"
	"os"
	"task-service/internal/api"
	task_service2 "task-service/internal/repository/task"
	"task-service/internal/usecase/task"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type diContainer struct {
	server     *api.TaskServiceServer
	useCase    api.TaskServiceUsecase
	repository task.TaskServiceRepository

	db *gorm.DB
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
		d.useCase = task.NewTaskServiceUseCase(d.Repository(ctx))
	}
	return d.useCase
}
func (d *diContainer) Repository(ctx context.Context) task.TaskServiceRepository {
	if d.repository == nil {
		d.repository = task_service2.NewTaskServiceRepository(d.DB(ctx))
	}
	return d.repository
}
func (d *diContainer) DB(ctx context.Context) *gorm.DB {
	if d.db == nil {
		db, err := gorm.Open(postgres.Open(os.Getenv("DB_URI")), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to db: %v", err)
		}
		d.db = db
	}
	return d.db
}
