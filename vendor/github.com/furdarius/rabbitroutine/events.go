package rabbitroutine

import "github.com/streadway/amqp"

// Retried is fired when connection retrying occurs.
// The event will be emitted only if the connection was not established.
// If connection was successfully established Dialed event emitted.
type Retried struct {
	ReconnectAttempt uint
	Error            error
}

// Dialed is fired when connection was successfully established.
type Dialed struct{}

// AMQPNotified is fired when AMQP error occurred.
type AMQPNotified struct {
	Error *amqp.Error
}
