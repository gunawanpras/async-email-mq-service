package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func (s *RabbitMQClient) Publish(ctx context.Context, exchange, routingKey string, payload any) error {
	ch, err := s.client.Connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Enable publisher confirms
	err = ch.Confirm(false)
	if err != nil {
		log.Fatalf("[publish] channel could not be put into confirm mode: %v", err)
	}

	log.Println("[publish] channel is in confirm mode...ok")

	// Set up a channel to receive publish confirmations
	confirms := ch.NotifyPublish(make(chan amqp091.Confirmation, 1))

	q, err := ch.QueueDeclare(
		routingKey, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Fatalf("[publish] failed to declare a queue: %v", err)
	}

	log.Println("[publish] declared queue...ok")

	err = ch.PublishWithContext(
		ctx,
		exchange,
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("[publish] failed to publish message: %w", err)
	}

	log.Println("[publish] published message...ok")

	select {
	case confirm := <-confirms:
		if confirm.Ack {
			log.Printf("[publish] acknowledge message...ok")
		} else {
			log.Printf("[publish] acknowledge message...failed")
		}
	case <-time.After(5 * time.Second):
		log.Printf("[publish] no confirmation received within the timeout period!")
	}

	return nil
}

func (s *RabbitMQClient) Subscribe(ctx context.Context, routingKey string, autoAck bool, handler func(amqp091.Delivery) error) (err error) {
	ch, err := s.client.Connection.Channel()
	if err != nil {
		return err
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		routingKey,
		"",
		autoAck, // auto-ack disabled for manual control
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("[Subscribe] failed to consume message:", err)
		return err
	}

	go func() {
		for msg := range msgs {
			err = handler(msg)
			if err != nil {
				msg.Nack(false, false)
				continue
			}

			msg.Ack(false)
		}
	}()

	return nil
}
