package setup

import (
	emailTaskService "github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
	emailTaskServiceV0 "github.com/gunawanpras/async-email-mq-service/internal/core/email-task/service"
)

type Service struct {
	EmailTaskService emailTaskService.Service
}

func NewService(repo Repository) Service {
	return Service{
		EmailTaskService: emailTaskServiceV0.New(emailTaskServiceV0.InitAttribute{
			Repo: emailTaskServiceV0.RepoAttribute{
				EmailTaskRepo: repo.EmailTaskRepo,
			},
		}),
	}
}
