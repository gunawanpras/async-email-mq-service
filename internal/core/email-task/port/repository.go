package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
)

type Repository interface {
	CreateEmailTask(ctx context.Context, emailTask domain.EmailTask) (id uuid.UUID, err error)
	GetEmailTaskByParams(ctx context.Context, args domain.EmailTask) (res domain.EmailTask, err error)
	GetUserByEmail(ctx context.Context, email string) (res domain.User, err error)
}
