package main

import (
	"fmt"
	"log"
	"os"

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

	// Declarar la cola
	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al declarar la cola: %s", err)
	}

	// Recibir mensajes
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al recibir mensajes: %s", err)
	}

	log.Println("[*] Esperando mensajes. Presiona CTRL+C para salir.")

	// Leer mensajes de la cola
	for msg := range msgs {
		log.Printf("[x] Recibido: %s", msg.Body)
	}
}
