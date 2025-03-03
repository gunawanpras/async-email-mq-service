package mq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type (
	RabbitMQ interface {
		Publish(ctx context.Context, exchange, routingKey string, payload any) error
		Subscribe(ctx context.Context, routingKey string, autoAck bool, handler func(amqp091.Delivery) error) (err error)
	}

	RabbitMQClient struct {
		client Client
	}

	InitAttribute struct {
		Client Client
	}

	Client struct {
		Connection *amqp091.Connection
	}
)
