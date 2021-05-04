package main

import (
	"fmt"
	"log"

	"github.com/jailtonjunior94/bookings/api/infrastructure/bus"
	"github.com/jailtonjunior94/bookings/notifier/handlers"
	"github.com/jailtonjunior94/bookings/notifier/infrastructure"
)

func main() {
	infrastructure.New()
	notification := handlers.NewNotificationHandler()

	rabbitMQ := bus.New(infrastructure.RabbitMQConnection)
	channel := rabbitMQ.GetChannel()
	defer channel.Close()

	messages, err := channel.Consume(infrastructure.RabbitMQQueue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			fmt.Printf("Recieved Message: %s\n", message.Body)
			go notification.SendMessage(message.Body)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - Waiting for messages")

	<-forever
}
