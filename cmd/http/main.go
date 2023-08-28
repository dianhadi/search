package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

const (
	serviceName = "search-http-service"
)

func main() {
	http.HandleFunc("/publish", publishHandler)
	http.ListenAndServe(":8009", nil)
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	connectionString := "amqp://admin:admin@rabbitmq:5672"
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"post-created", // queue name
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	body := "Hello, RabbitMQ!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Println("Message sent:", body)
}
