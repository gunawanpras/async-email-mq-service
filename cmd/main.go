package main

import (
	"github.com/gunawanpras/async-email-mq-service/config"
	"github.com/gunawanpras/async-email-mq-service/internal/setup"
)

func main() {
	// init config
	config.Init(config.WithConfigFile("config"), config.WithConfigType("yaml"))
	conf := config.Get()

	// init external services
	externalService := setup.InitExternalServices(conf)
	defer externalService.Postgres.Close()
	defer externalService.RabbitMq.Close()

	// init core services
}
