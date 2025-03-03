package v0

import (
	"context"
	"errors"
	"log"

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
		if err.Error() != constant.DataNotFound {
			return res, err
		}
	}

	if emailTask.ID != uuid.Nil {
		return emailTask, errors.New(constant.EmailTaskAlreadyExist)
	}

	now := timeutil.TimeHelper.Now()
	task := domain.EmailTask{
		SenderEmail:    args.SenderEmail,
		RecipientEmail: args.RecipientEmail,
		Subject:        args.Subject,
		Body:           args.Body,
		Status:         constant.EmailTaskStatusPending,
		CreatedAt:      now,
		CreatedBy:      constant.SYSTEM,
	}

	id, err := service.repo.EmailTaskRepo.CreateEmailTask(ctx, task)
	if err != nil {
		return res, err
	}

	task.ID = id
	err = service.mq.EmailTaskMQ.PublishEmailTask(ctx, task)
	if err != nil {
		return res, err
	}

	return task, nil
}

func (service *EmailTaskService) ProcessEmailTask(ctx context.Context) (err error) {
	message, err := service.mq.EmailTaskMQ.SubscribeEmailTask(ctx)
	if err != nil {
		return err
	}

	emailTask, err := service.repo.EmailTaskRepo.GetEmailTaskByID(ctx, message.ID)
	if err != nil {
		if err.Error() == constant.DataNotFound {
			return errors.New(constant.EmailTaskNotFound)
		}
		return err
	}

	if emailTask.Status == constant.EmailTaskStatusPending {
		// TODO: Send email
		log.Println("[email-task] process email-task ID: " + emailTask.ID.String() + "...ok")
	}

	return nil
}
