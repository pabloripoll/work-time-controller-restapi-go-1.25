package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	URL string
}

type Connection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewConnection(config Config) (*Connection, error) {
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	log.Println("✅ RabbitMQ connection established")

	return &Connection{
		conn:    conn,
		channel: channel,
	}, nil
}

func (c *Connection) Close() error {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Connection) Channel() *amqp.Channel {
	return c.channel
}
