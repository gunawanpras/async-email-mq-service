package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
	"github.com/gunawanpras/async-email-mq-service/pkg/util/constant"
	"github.com/gunawanpras/async-email-mq-service/pkg/util/timeutil"
)

func (service *EmailTaskService) CreateEmailTask(ctx context.Context, args domain.EmailTask) (res domain.EmailTask, err error) {
	_, err = service.repo.EmailTaskRepo.GetUserByEmail(ctx, args.SenderEmail)
	if err != nil {
		if err.Error() == constant.DataNotFound {
			return res, errors.New(constant.SenderEmailNotFound)
		}
		return res, err
	}

	emailTask, err := service.repo.EmailTaskRepo.GetEmailTaskByParams(ctx, args)
	if err != nil {
		return res, err
	}

	if emailTask.ID != uuid.Nil {
		return emailTask, errors.New(constant.EmailTaskAlreadyExist)
	}

	now := timeutil.TimeHelper.Now()

	id, err := service.repo.EmailTaskRepo.CreateEmailTask(ctx, domain.EmailTask{
		SenderEmail:    args.SenderEmail,
		RecipientEmail: args.RecipientEmail,
		Subject:        args.Subject,
		Body:           args.Body,
		Status:         constant.EmailTaskStatusPending,
		CreatedAt:      now,
		CreatedBy:      constant.SYSTEM,
	})

	if err != nil {
		return res, err
	}

	res = domain.EmailTask{
		ID:             id,
		SenderEmail:    args.SenderEmail,
		RecipientEmail: args.RecipientEmail,
		Subject:        args.Subject,
		Body:           args.Body,
		Status:         constant.EmailTaskStatusPending,
		CreatedAt:      now,
		CreatedBy:      constant.SYSTEM,
	}

	return res, nil
}
