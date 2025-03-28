package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	if rabbitmqHost == "" {
		rabbitmqHost = "localhost"
	}

	conn, err := amqp091.Dial("amqp://guest:guest@" + rabbitmqHost + ":5672/")
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir un canal: %s", err)
	}
	defer ch.Close()

	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		queueName = "go_queue"
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola: %s", err)
	}

	for i := 1; i <= 5; i++ {
		body := fmt.Sprintf("Mensaje %d desde Go Producer!", i)
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		if err != nil {
			log.Fatalf("Error al publicar el mensaje: %s", err)
		}
		log.Printf("[x] Enviado: %s", body)
		time.Sleep(2 * time.Second)
	}
}
