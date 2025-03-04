package setup

import (
	"github.com/gunawanpras/async-email-mq-service/config"
	setupClient "github.com/gunawanpras/async-email-mq-service/internal/setup/client"
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
	pg := setupClient.InitPostgres(conf)
	mq := setupClient.InitMq(conf)

	return &ExternalServices{
		Postgres: pg,
		RabbitMq: mq,
	}
}

func InitCoreServices(externalService *ExternalServices) *CoreServices {
	repository := NewRepository(externalService.Postgres)
	event := NewMessageQueue(externalService.RabbitMq)
	service := NewService(repository, event)
	handler := NewHandler(service)

	return &CoreServices{
		Handler: *handler,
	}
}
