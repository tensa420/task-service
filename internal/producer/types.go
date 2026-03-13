package producer

import "context"

type LogProducer interface {
	SendMessage(ctx context.Context, topic string, userUUID, msg string) error
}
