package kinkajou

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type HandlerFunc func(msg *amqp.Delivery)

func (app *App) OnEvent(routingKey string, handler HandlerFunc) {
	app.handlers[routingKey] = handler
}

func (app *App) Dispatch(msg *amqp.Delivery) {
	if handler, ok := app.handlers[msg.RoutingKey]; ok {
		handler(msg)
	} else {
		app.defaultHandler(msg)
	}
}
