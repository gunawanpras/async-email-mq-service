package setup

import handler "github.com/gunawanpras/async-email-mq-service/internal/adapter/http/handler/email-task"

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
