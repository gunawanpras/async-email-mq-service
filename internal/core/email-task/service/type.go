package v0

import (
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
)

type (
	RepoAttribute struct {
		EmailTaskRepo port.Repository
	}

	Config struct{}

	EmailTaskService struct {
		repo   RepoAttribute
		config Config
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config Config
	}
)
