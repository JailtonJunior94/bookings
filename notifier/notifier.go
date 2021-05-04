package main

import (
	"fmt"
	"log"

	"github.com/jailtonjunior94/bookings/api/infrastructure/bus"
	"github.com/jailtonjunior94/bookings/notifier/infrastructure"
)

func main() {
	infrastructure.New()

	rabbitMQ := bus.New(infrastructure.RabbitMQConnection)

	messages, err := rabbitMQ.GetChannel().Consume(infrastructure.RabbitMQQueue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
