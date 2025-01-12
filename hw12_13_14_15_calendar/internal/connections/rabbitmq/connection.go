package rabbitmq

import (
	"github.com/streadway/amqp"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// RabbitMQ describes the RabbitMQ connection.
type RabbitMQ struct {
	// conn describes the RabbitMQ connection.
	conn *amqp.Connection
	// channel describes the RabbitMQ channel.
	channel *amqp.Channel
	// queue describes the RabbitMQ queue.
	queue amqp.Queue
}

// NewRabbitMQ creates and returns a new RabbitMQ connection.
func NewRabbitMQ(url, queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

// Publish publishes the message to the RabbitMQ queue.
func (r *RabbitMQ) PublishJSON(exchange, key string, body []byte) error {
	return r.channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// Consume consumes the messages from the RabbitMQ queue.
func (r *RabbitMQ) Consume() (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		r.queue.Name,
		"",
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
