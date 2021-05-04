package bus

import (
	"log"

	"github.com/jailtonjunior94/bookings/api/domain/interfaces"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func New(connectionString string) interfaces.IRabbitMQ {
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &RabbitMQ{Connection: connection, Channel: channel}
}

func (r *RabbitMQ) SendMessage(queue, message string) error {
	_, err := r.Channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	if err := r.Channel.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) GetChannel() *amqp.Channel {
	return r.Channel
}
