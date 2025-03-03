package emailtask

import (
	"fmt"
	"log"

	"github.com/gunawanpras/async-email-mq-service/internal/core/email-task/port"
)

func NewEmailTaskMessageQueue(attr InitAttribute) port.MessageQueue {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	mq := &EmailTaskMessageQueue{
		rabbitmq: attr.RabbitMQClient,
	}

	return mq
}

func (init InitAttribute) validate() error {
	if !init.RabbitMQClient.validate() {
		return fmt.Errorf("missing mq client : %+v", init.RabbitMQClient)
	}

	return nil
}

func (client RabbitMQClient) validate() bool {
	return client.RabbitMQClient != nil
}
