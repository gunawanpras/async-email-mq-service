package port

import (
	"context"

	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
)

type MessageQueue interface {
	PublishEmailTask(ctx context.Context, task domain.EmailTask) (err error)
	SubscribeEmailTask(ctx context.Context) (res domain.EmailTask, err error)
}
