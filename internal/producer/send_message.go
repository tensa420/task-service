package producer

import (
	"context"
	"log"
	"task-service/internal/entity/events"
	protoEvents "task-service/pkg/events"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (p *LogsProducer) SendMessage(ctx context.Context, taskLog events.TaskLog) error {
	protoLog := convertEntityTaskLogToProto(taskLog)
	marshalledLog, err := proto.Marshal(&protoLog)
	if err != nil {
		return err
	}
	partition, offset, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(taskLog.LogUUID.String()),
		Value: sarama.ByteEncoder(marshalledLog),
	})
	log.Printf("Partition: %d, Offset: %d", partition, offset)
	if err != nil {
		return err
	}
	return nil
}

func convertEntityTaskLogToProto(log events.TaskLog) protoEvents.TaskLog {
	return protoEvents.TaskLog{
		TaskUUID:  log.TaskUUID.String(),
		UserUUID:  log.UserUUID.String(),
		LogType:   convertEntityLogTypeToProto(log.Type),
		CreatedAt: timestamppb.New(log.Created_at),
		LogUUID:   log.LogUUID.String(),
	}
}
func convertEntityLogTypeToProto(logType events.LogType) protoEvents.LogType {
	switch logType {
	case events.LogTypeFinish:
		return protoEvents.LogType_LOG_TYPE_FINISH
	case events.LogTypeCreate:
		return protoEvents.LogType_LOG_TYPE_CREATE
	default:
		return protoEvents.LogType_LOG_TYPE_DELETE
	}
}
