package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/async-email-mq-service/internal/adapter/email-task/http/dto"
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
	"github.com/gunawanpras/async-email-mq-service/pkg/response"
	"github.com/gunawanpras/async-email-mq-service/pkg/util/constant"
	"github.com/gunawanpras/async-email-mq-service/pkg/validator"
)

func (handler *EmailTaskHandler) CreateEmailTask(c *fiber.Ctx) error {
	var req dto.CreateEmailTaskRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	emailTask := domain.EmailTask{
		SenderEmail:    req.SenderEmail,
		RecipientEmail: req.RecipientEmail,
		Subject:        req.Subject,
		Body:           req.Body,
	}

	resp, err := handler.service.EmailTaskService.CreateEmailTask(ctx, emailTask)
	if err != nil {
		return response.Error(c, constant.EmailTaskCreateFailed, err, constant.EmailTaskHttpStatusMappings)
	}

	respData := dto.CreateEmailTaskResponse{
		ID: resp.ID,
	}

	return response.OK(c, constant.EmailTaskCreateSuccess, respData, constant.EmailTaskHttpStatusMappings)
}
