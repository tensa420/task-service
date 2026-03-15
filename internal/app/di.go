package app

import (
	"context"
	"log"
	"os"
	"strings"
	"task-service/internal/api"
	"task-service/internal/producer"
	task_service2 "task-service/internal/repository/task"
	"task-service/internal/usecase/task"

	"github.com/IBM/sarama"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type diContainer struct {
	server      *api.TaskServiceServer
	useCase     api.TaskServiceUsecase
	repository  task.TaskRepository
	producer    sarama.SyncProducer
	logProducer producer.LogProducer
	db          *gorm.DB
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
		d.useCase = task.NewTaskServiceUseCase(d.Repository(ctx), d.LogProducer(ctx))
	}
	return d.useCase
}

func (d *diContainer) Repository(ctx context.Context) task.TaskRepository {
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

func (d *diContainer) Producer(ctx context.Context) sarama.SyncProducer {
	if d.producer == nil {
		brokersEnv := os.Getenv("KAFKA_BROKERS")
		if brokersEnv == "" {
			brokersEnv = "localhost:9092"
		}
		brokers := strings.Split(brokersEnv, ",")

		cfg := sarama.NewConfig()
		cfg.Producer.Return.Successes = true

		producer, err := sarama.NewSyncProducer(brokers, cfg)
		if err != nil {
			log.Fatalf("failed to create kafka producer: %v", err)
		}
		d.producer = producer
	}
	return d.producer
}

func (d *diContainer) LogProducer(ctx context.Context) producer.LogProducer {
	if d.logProducer == nil {
		topic := os.Getenv("KAFKA_TOPIC")
		if topic == "" {
			topic = "task_logs"
		}
		d.logProducer = producer.NewOrderPaidProducer(d.Producer(ctx), topic)
	}
	return d.logProducer
}
