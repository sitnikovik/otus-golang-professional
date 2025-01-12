package rabbitmq

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// RabbitMQ describes the RabbitMQ connection.
type RabbitMQ struct {
	// conn describes the RabbitMQ connection.
	conn *amqp.Connection
	// channel describes the RabbitMQ channel.
	channel *amqp.Channel
	// queue is the RabbitMQ queue that is declared on channel connected.
	queue amqp.Queue
	// exchangeName is the RabbitMQ exchange name that is declared on channel connected.
	exchangeName string
	// exchangeType is the RabbitMQ exchange type that is declared on channel connected.
	exchangeType string
}

// NewRabbitMQ creates and returns a new RabbitMQ connection.
func NewRabbitMQ(dsn, queueName, exchangeName, exchangeType string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	rmq := &RabbitMQ{
		conn:    conn,
		channel: ch,
	}

	// Declare exchange
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange \"%s\": %v", exchangeName, err)
	}
	rmq.exchangeName = exchangeName
	rmq.exchangeType = exchangeType

	// Declare queue
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue \"%s\": %v", queueName, err)
	}
	rmq.queue = q

	err = ch.QueueBind(
		queueName,
		queueName,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to bind queue \"%s\" to exchange \"%s\": %v",
			queueName,
			exchangeName,
			err,
		)
	}

	return rmq, nil
}

// NewDSN creates and returns the RabbitMQ DSN.
func NewDSN(host, port, login, pwd string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s", login, pwd, host, port)
}

// Publish publishes the message to the RabbitMQ queue.
func (r *RabbitMQ) PublishJSON(_ context.Context, key string, body []byte) error {
	return r.channel.Publish(
		r.exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// Consume consumes the messages from the RabbitMQ queue.
func (r *RabbitMQ) Consume(_ context.Context, consumer string) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		r.queue.Name,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)
}

// Close closes the RabbitMQ connection.
func (r *RabbitMQ) Close() {
	if err := r.channel.Close(); err != nil {
		logger.Alertf("Failed to close channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		logger.Alertf("Failed to close connection: %v", err)
	}
}
