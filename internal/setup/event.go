package setup

import (
	mq "github.com/gunawanpras/async-email-mq-service/internal/adapter/mq/rabbitmq"
	emailTaskMQ "github.com/gunawanpras/async-email-mq-service/internal/adapter/mq/rabbitmq/email-task"
	emailTaskMQPort "github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
	"github.com/rabbitmq/amqp091-go"
)

type MessageQueue struct {
	EmailTaskMQ emailTaskMQPort.MessageQueue
}

func NewMessageQueue(conn *amqp091.Connection) MessageQueue {
	rabbitMQClient := mq.NewRabbitMQClient(mq.InitAttribute{
		Client: mq.Client{
			Connection: conn,
		},
	})

	emailTaskMQ := emailTaskMQ.NewEmailTaskMessageQueue(emailTaskMQ.InitAttribute{
		RabbitMQClient: emailTaskMQ.RabbitMQClient{
			RabbitMQClient: rabbitMQClient,
		},
	})

	return MessageQueue{
		EmailTaskMQ: emailTaskMQ,
	}
}
