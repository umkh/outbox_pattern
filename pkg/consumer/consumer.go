package consumer

import (
	"context"
	"log"

	"github.com/furdarius/rabbitroutine"
	"github.com/streadway/amqp"
)

type Params struct {
	Connector *rabbitroutine.Connector
	Exchange  string
	Binding   string
	QueueName string
}

type Consumer struct {
	conn      *rabbitroutine.Connector
	exchange  string
	binding   string
	queueName string
	delivery  func(msg amqp.Delivery)
}

func New(p *Params) *Consumer {
	return &Consumer{
		exchange:  p.Exchange,
		binding:   p.Binding,
		queueName: p.QueueName,
		conn:      p.Connector,
	}
}

func (c *Consumer) Declare(ctx context.Context, ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		c.exchange, // name
		"topic",    // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Printf("failed to declare exchange %v: %v", c.exchange, err)

		return err
	}

	_, err = ch.QueueDeclare(
		c.queueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,
	)
	if err != nil {
		log.Printf("failed to declare queue %v: %v", c.queueName, err)

		return err
	}

	err = ch.QueueBind(
		c.queueName, // queue name
		c.binding,   // routing key
		c.exchange,  // exchange
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Printf("failed to bind queue %v: %v", c.queueName, err)

		return err
	}

	return nil
}

func (c *Consumer) Consume(ctx context.Context, ch *amqp.Channel) error {
	defer log.Println("consume method finished")

	err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Printf("failed to set qos: %v", err)

		return err
	}

	msgs, err := ch.Consume(
		c.queueName, // queue
		"",          // consumer name
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		log.Printf("failed to consume %v: %v", c.queueName, err)

		return err
	}

	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				return amqp.ErrClosed
			}

			c.delivery(msg)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *Consumer) StartConsumer(ctx context.Context, delivery func(msg amqp.Delivery)) error {
	c.delivery = delivery
	return c.conn.StartConsumer(ctx, c)
}
