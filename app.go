package kinkajou

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type App struct {
	handlers       map[string]HandlerFunc
	defaultHandler HandlerFunc
}

func New() *App {
	return &App{
		handlers: make(map[string]HandlerFunc),
		defaultHandler: func(msg *amqp.Delivery) {
			log.Printf("Received message by default handler: %s", msg.Body)

		},
	}
}
