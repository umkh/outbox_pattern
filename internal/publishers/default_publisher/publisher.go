package default_publisher

import (
	"context"

	"github.com/streadway/amqp"
	"github.com/umkh/outbox_pattern/pkg/publisher"

	"github.com/furdarius/rabbitroutine"
)

type Publisher struct {
	pub *publisher.Publisher
}

func New(cfg Config, rb *rabbitroutine.Connector) *Publisher {
	publisher := publisher.New(&publisher.Params{
		Connector: rb,
		Exchange:  cfg.Exchange,
		Binding:   cfg.Binding,
	})

	return &Publisher{
		pub: publisher,
	}
}

func (df *Publisher) Push(ctx context.Context) error {
	return df.pub.Publish(ctx, amqp.Publishing{
		MessageId:    "12345",
		DeliveryMode: 2,
		Body:         []byte("test"),
	})
}
