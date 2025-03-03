package mq

import (
	"fmt"
	"log"
)

func NewRabbitMQClient(attr InitAttribute) RabbitMQ {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	mq := &RabbitMQClient{
		client: attr.Client,
	}

	return mq
}

func (init InitAttribute) validate() error {
	if !init.Client.validate() {
		return fmt.Errorf("missing mq client : %+v", init.Client)
	}

	return nil
}

func (client Client) validate() bool {
	return client.Connection != nil
}
