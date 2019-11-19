package externals

import (
	"github.com/streadway/amqp"
)

//AMQPConnection represents the needed functionality from a amqp.Connection
type AMQPConnection interface {
	Channel() (*amqp.Channel, error)
	Close() error
}
