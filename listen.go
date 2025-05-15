package kinkajou

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (app *App) Listen(messages <-chan amqp.Delivery) error {
	for msg := range messages {
		app.Dispatch(&msg)
	}
	return nil
}
