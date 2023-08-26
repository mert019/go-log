package logdestinations

import (
	"context"
	"encoding/json"

	"github.com/mert019/go-log/gologcore"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQLoggerConfiguration struct {
	Url       string
	QueueName string
}

func NewRabbitMqLogger(configuration RabbitMQLoggerConfiguration) (gologcore.ILogDestination, error) {

	connection, err := amqp.Dial(configuration.Url)
	if err != nil {
		return nil, &gologcore.LogDestinationConnectionError{
			Destination:     "RabbitMQ",
			ConnectionError: err,
		}
	}

	ch, err := connection.Channel()
	if err != nil {
		return nil, &gologcore.LogDestinationConnectionError{
			Destination:     "RabbitMQ",
			ConnectionError: err,
		}
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		configuration.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &RabbitMQLogger{
		configuration: configuration,
		connection:    connection,
		queue:         queue,
	}, nil
}

type RabbitMQLogger struct {
	configuration RabbitMQLoggerConfiguration
	connection    *amqp.Connection
	queue         amqp.Queue
}

func (rabbitMQLogger *RabbitMQLogger) Log(log gologcore.Log) error {

	channel, err := rabbitMQLogger.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	body, err := json.Marshal(log)
	if err != nil {
		return err
	}

	return channel.PublishWithContext(
		context.Background(),
		"",
		rabbitMQLogger.queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		})
}

func (rabbitMQLogger *RabbitMQLogger) Close() error {
	return rabbitMQLogger.connection.Close()
}
