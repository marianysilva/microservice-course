package shared

import (
	"fmt"

	"github.com/sumelms/microservice-course/pkg/config"
	"github.com/sumelms/microservice-course/pkg/logger"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPConnection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Config     *config.AMQPClient
	logger     logger.Logger
}

type AMQPQueue struct {
	Name    string
	Queue   *amqp.Queue
	Channel *amqp.Channel
}

func NewAMQPConnection(cfg *config.AMQPClient, logger logger.Logger) (*AMQPConnection, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid server config")
	}

	conn, err := amqp.Dial(cfg.Host)
	if err != nil {
		logger.Log("msg", "unable to connect with RabbitMQ server: ", "error", err)
		return nil, err
	}
	logger.Log("msg", "Connected with RabbitMQ server! ", "host", cfg.Host)

	ch, err := conn.Channel()
	if err != nil {
		logger.Log("msg", "unable to create RabbitMQ channel: ", "error", err)
		return nil, err
	}

	logger.Log("msg", "Using RabbitMQ channel!")

	return &AMQPConnection{
		Connection: conn,
		Channel:    ch,
		Config:     cfg,
		logger:     logger,
	}, nil
}

func (c AMQPConnection) NewAMQPQueue(queueName string) (*AMQPQueue, error) {
	if c.Channel == nil {
		c.logger.Log("msg", "unable to create RabbitMQ channel ")
		return nil, fmt.Errorf("unable to use RabbitMQ channel")
	}

	queue, err := c.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		c.logger.Log("msg", "unable to create RabbitMQ queue: ", "error", err)
		return nil, err
	}

	return &AMQPQueue{
		Name:    queueName,
		Queue:   &queue,
		Channel: c.Channel,
	}, nil
}

func (queue AMQPQueue) Publish(jsonData []byte) error {
	err := queue.Channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	)

	return err
}
