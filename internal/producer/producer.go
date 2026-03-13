package producer

import "github.com/IBM/sarama"

type LogsProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewOrderPaidProducer(producer sarama.SyncProducer, topic string) *LogsProducer {
	return &LogsProducer{
		producer: producer,
		topic:    topic,
	}
}
