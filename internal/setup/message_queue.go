package setup

import (
	"log"

	"github.com/gunawanpras/async-email-mq-service/config"
	"github.com/rabbitmq/amqp091-go"
)

func InitMq(conf *config.Config) *amqp091.Connection {
	conn, err := amqp091.Dial(conf.RabbitMQ.Primary.ConnString)
	if err != nil {
		log.Panic("failed to connect to RabbitMQ service:", err)
	}

	return conn
}
