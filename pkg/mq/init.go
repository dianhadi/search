package mq

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

type HandlerConsumer func(ctx context.Context, body []byte) error

func New(host string, port int, username, password string) (*MQ, error) {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d", username, password, host, port)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	mq := MQ{
		conn:    conn,
		channel: ch,
	}

	return &mq, err
}

func (m MQ) Close() {
	m.conn.Close()
	m.channel.Close()
}
