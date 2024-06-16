package main

import (
	"fmt"
	"log"

	"github.com/henrique77/rabbitmq_producer_consumer/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Canais para mensagens
	msgsEven := make(chan amqp.Delivery)
	msgsOdd := make(chan amqp.Delivery)

	// Consome mensagens da fila "even_numbers"
	go func() {
		err := rabbitmq.Consume(ch, msgsEven, "even_numbers", "go-consumer-even")
		if err != nil {
			log.Fatalf("Failed to consume from even_numbers queue: %v", err)
		}
	}()

	// Consome mensagens da fila "odd_numbers"
	go func() {
		err := rabbitmq.Consume(ch, msgsOdd, "odd_numbers", "go-consumer-old")
		if err != nil {
			log.Fatalf("Failed to consume from odd_numbers queue: %v", err)
		}
	}()

	// Processa mensagens de ambas as filas
	for {
		select {
		case msg := <-msgsEven:
			fmt.Printf("Received even number: %s\n", string(msg.Body))
			msg.Ack(false)
		case msg := <-msgsOdd:
			fmt.Printf("Received odd number: %s\n", string(msg.Body))
			msg.Ack(false)
		}
	}
}
