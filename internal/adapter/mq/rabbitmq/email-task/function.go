package emailtask

import (
	"context"
	"encoding/json"

	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/domain"
	"github.com/gunawanpras/async-email-mq-service/pkg/util/constant"
	"github.com/rabbitmq/amqp091-go"
)

func (r *EmailTaskMessageQueue) PublishEmailTask(ctx context.Context, task domain.EmailTask) (err error) {
	exchange := ""
	routingKey := constant.EmailTaskQueueName

	err = r.rabbitmq.RabbitMQClient.Publish(ctx, exchange, routingKey, task)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmailTaskMessageQueue) SubscribeEmailTask(ctx context.Context) (res domain.EmailTask, err error) {
	routingKey := constant.EmailTaskQueueName

	resultChan := make(chan domain.EmailTask)
	defer close(resultChan)

	if err = r.rabbitmq.RabbitMQClient.Subscribe(ctx, routingKey, false, func(msg amqp091.Delivery) error {
		var emailTask domain.EmailTask

		if err := json.Unmarshal(msg.Body, &emailTask); err != nil {
			return err
		}

		resultChan <- emailTask
		return nil

	}); err != nil {
		return res, err
	}

	res = <-resultChan

	return res, nil
}
