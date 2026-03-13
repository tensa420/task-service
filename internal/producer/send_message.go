package producer

import (
	"context"
	"log"
	"task-service/internal/entity"
	"task-service/pkg/task_service"

	"github.com/IBM/sarama"
)

func (p *LogsProducer) SendMessage(ctx context.Context, topic string, userUUID, msg string) error {
	partition, offset, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(userUUID),
		Value: sarama.ByteEncoder(msg),
	})
	log.Printf("Partition:%d,Offset:%d", partition, offset)
	if err != nil {
		return err
	}
	return nil
}
func convertEntityStatusToProto(status entity.TaskStatus) task_service.TaskStatus {
	switch status {
	case entity.TaskStatusNew:
		return task_service.TaskStatus_TASK_STATUS_NEW
	default:
		return task_service.TaskStatus_TASK_STATUS_FINISHED
	}
}
