package setup

import (
	"github.com/gunawanpras/async-email-mq-service/config"
	"github.com/jmoiron/sqlx"
	"github.com/rabbitmq/amqp091-go"
)

type ExternalServices struct {
	Postgres *sqlx.DB
	RabbitMq *amqp091.Connection
}

type CoreServices struct {
	Handler Handler
}

func InitExternalServices(conf *config.Config) *ExternalServices {
	pg := InitPostgres(conf)
	mq := InitMq(conf)

	return &ExternalServices{
		Postgres: pg,
		RabbitMq: mq,
	}
}
