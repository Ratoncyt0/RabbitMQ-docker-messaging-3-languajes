package main

import (
	"fmt"
	"log"
	"os"

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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error al consumir mensajes: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[x] Recibido: %s", d.Body)
		}
	}()

	fmt.Println("[*] Esperando mensajes. Presiona CTRL+C para salir.")
	<-forever
}
