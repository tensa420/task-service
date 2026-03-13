package producer

import (
	"context"
	"task-service/internal/entity/events"
)

type LogProducer interface {
	SendMessage(ctx context.Context, taskLog events.TaskLog) error
}
