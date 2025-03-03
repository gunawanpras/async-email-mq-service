package port

import (
	"context"

	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
)

type Service interface {
	CreateEmailTask(ctx context.Context, emailTask domain.EmailTask) (res domain.EmailTask, err error)
	ProcessEmailTask(ctx context.Context) (err error)
}
