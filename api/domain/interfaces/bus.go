package interfaces

import "github.com/streadway/amqp"

type IRabbitMQ interface {
	SendMessage(queue, message string) error
	GetChannel() *amqp.Channel
}
