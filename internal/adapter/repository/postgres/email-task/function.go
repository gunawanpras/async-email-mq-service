package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
	"github.com/gunawanpras/async-email-mq-service/pkg/util/constant"
	"github.com/gunawanpras/async-email-mq-service/pkg/util/uuidutil"
)

// Create new email task and return the generated ID
func (repo *EmailTaskRepository) CreateEmailTask(ctx context.Context, args domain.EmailTask) (id uuid.UUID, err error) {
	args.ID = uuidutil.UUIDHelper.New()

	repo.prepareCreateEmailTask()
	_, err = repo.statement.CreateEmailTask.ExecContext(ctx, args.ID, args.SenderEmail, args.RecipientEmail, args.Subject, args.Body, args.Status, args.CreatedAt, args.CreatedBy)
	if err != nil {
		return uuid.Nil, err
	}
	return args.ID, nil
}

func (repo *EmailTaskRepository) GetEmailTaskByParams(ctx context.Context, args domain.EmailTask) (res domain.EmailTask, err error) {
	var emailTask EmailTask

	repo.prepareGetEmailTaskByParams()
	err = repo.statement.GetEmailTaskByParams.QueryRowxContext(ctx, args.SenderEmail, args.RecipientEmail, args.Subject, args.Body).StructScan(&emailTask)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !emailTask.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return emailTask.ToModel(), nil
}

func (repo *EmailTaskRepository) GetEmailTaskByID(ctx context.Context, id uuid.UUID) (res domain.EmailTask, err error) {
	var emailTask EmailTask

	repo.prepareGetEmailTaskByID()
	err = repo.statement.GetEmailTaskByID.QueryRowxContext(ctx, id).StructScan(&emailTask)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !emailTask.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return emailTask.ToModel(), nil
}

func (repo *EmailTaskRepository) GetUserByEmail(ctx context.Context, email string) (res domain.User, err error) {
	var user User

	repo.prepareGetUserByEmail()
	err = repo.statement.GetUserByEmail.QueryRowxContext(ctx, email).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !user.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return user.ToModel(), nil
}
