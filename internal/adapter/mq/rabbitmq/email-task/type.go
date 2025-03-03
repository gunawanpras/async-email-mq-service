package emailtask

import mq "github.com/gunawanpras/async-email-mq-service/internal/adapter/mq/rabbitmq"

type (
	RabbitMQClient struct {
		RabbitMQClient mq.RabbitMQ
	}

	EmailTaskMessageQueue struct {
		rabbitmq RabbitMQClient
	}

	InitAttribute struct {
		RabbitMQClient RabbitMQClient
	}
)
