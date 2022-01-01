package default_worker

import (
	"context"
	"fmt"
	"log"

	"github.com/umkh/outbox_pattern/pkg/consumer"

	"github.com/furdarius/rabbitroutine"
	"github.com/streadway/amqp"
)

type Worker struct {
	consumer *consumer.Consumer
	// outbox   *outbox.Outbox
}

func New(cfg Config, rb *rabbitroutine.Connector) *Worker {
	consumer := consumer.New(&consumer.Params{
		Connector: rb,
		Exchange:  cfg.Exchange,
		Binding:   cfg.Binding,
		QueueName: cfg.QueueName,
	})

	return &Worker{
		consumer: consumer,
	}
}

func (w *Worker) handle(msg amqp.Delivery) {
	content := string(msg.Body)

	fmt.Println("New message:", content)

	err := msg.Ack(false)
	if err != nil {
		log.Printf("failed to Ack message: %v", err)
	}
}

func (w *Worker) Run(ctx context.Context) error {
	return w.consumer.StartConsumer(ctx, w.handle)
}
