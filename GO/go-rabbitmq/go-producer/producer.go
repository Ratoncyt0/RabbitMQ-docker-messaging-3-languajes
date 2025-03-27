package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	// Obtener variables de entorno
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPass := os.Getenv("RABBITMQ_PASS")
	queueName := os.Getenv("QUEUE_NAME")

	// Conectar a RabbitMQ
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPass, rabbitmqHost, rabbitmqPort)
	conn, err := amqp091.Dial(connStr)
	if err != nil {
		log.Fatalf("Error de conexión a RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir un canal: %s", err)
	}
	defer ch.Close()

	// Declarar una cola
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al declarar la cola: %s", err)
	}

	// Enviar mensajes cada 2 segundos
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Mensaje %d desde Go Producer!", i)
		err = ch.Publish("", queueName, false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
		if err != nil {
			log.Fatalf("Error al enviar mensaje: %s", err)
		}

		log.Printf("[x] Enviado: %s", message)
		time.Sleep(2 * time.Second)
	}
}
