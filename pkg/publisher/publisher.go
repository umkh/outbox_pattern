package publisher

import (
	"context"

	"github.com/furdarius/rabbitroutine"
	"github.com/streadway/amqp"
)

type Params struct {
	Connector *rabbitroutine.Connector
	Exchange  string
	Binding   string
}

type Publisher struct {
	pub      *rabbitroutine.RetryPublisher
	exchange string
	binding  string
}

func New(p *Params) *Publisher {
	pool := rabbitroutine.NewPool(p.Connector)
	ensurePub := rabbitroutine.NewEnsurePublisher(pool)
	pub := rabbitroutine.NewRetryPublisher(ensurePub)

	return &Publisher{
		pub:      pub,
		exchange: p.Exchange,
		binding:  p.Binding,
	}
}

func (p *Publisher) Publish(ctx context.Context, msg amqp.Publishing) error {
	return p.pub.Publish(ctx, p.exchange, p.binding, msg)
}
