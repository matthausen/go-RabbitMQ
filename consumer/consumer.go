package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue", // queue name
		"",          //consumer
		true,        // auto-ack
		false,       //exclusive
		false,       // no-local
		false,       // no-wait
		nil,         //args
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to our rabbitMQ instance")
	fmt.Println(" [+] - waiting for messages")

	<-forever
}
