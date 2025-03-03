package v0

import (
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
)

type (
	RepoAttribute struct {
		EmailTaskRepo port.Repository
	}

	MQAttribute struct {
		EmailTaskMQ port.MessageQueue
	}

	Config struct{}

	EmailTaskService struct {
		repo   RepoAttribute
		mq     MQAttribute
		config Config
	}

	InitAttribute struct {
		Repo   RepoAttribute
		MQ     MQAttribute
		Config Config
	}
)
