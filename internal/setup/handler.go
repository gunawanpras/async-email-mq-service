package setup

import "github.com/gunawanpras/async-email-mq-service/internal/adapter/email-task/http/handler"

type Handler struct {
	EmailTaskHandler handler.Handler
}

func NewHandler(service Service) *Handler {
	return &Handler{
		EmailTaskHandler: handler.New(handler.InitAttribute{
			Service: handler.ServiceAttribute{
				EmailTaskService: service.EmailTaskService,
			},
		}),
	}
}
