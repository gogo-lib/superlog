package superlog

import (
	"context"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

type kafkaWriter struct {
	writer        *kafka.Writer
	batchSize     int
	messages      []kafka.Message
	messagesMutex sync.Mutex
}

func newKafkaWriter(writerConn *kafka.Writer, batchSize int) *kafkaWriter {
	kafkaWriter := &kafkaWriter{
		writer:    writerConn,
		batchSize: batchSize,
		messages:  make([]kafka.Message, 0, batchSize),
	}

	// messages is cached, so we need handle graceful shutdown

	return kafkaWriter
}

// Write implement io.Writer
func (w *kafkaWriter) Write(p []byte) (n int, err error) {
	w.messagesMutex.Lock()
	defer w.messagesMutex.Unlock()

	w.messages = append(w.messages, kafka.Message{
		Value: p,
	})

	// check batched message size
	if len(w.messages) == w.batchSize {
		// ref to batched messages
		batchedMessage := w.messages
		// update batched messages ref
		w.messages = make([]kafka.Message, 0, w.batchSize)
		// write batched messages to kafka
		go func() {
			err := w.writer.WriteMessages(context.Background(), batchedMessage...)
			if err != nil {
				// handle write kafka error
				log.Println(err)
			}
		}()
		return 1, nil
	}

	return 1, nil
}
