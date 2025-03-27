package main

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go" // Importar el paquete amqp091-go correctamente
)

func main() {
	// Conectar a RabbitMQ
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/") // Usar amqp091 para referirse al paquete
	if err != nil {
		log.Fatalf("Error de conexión: %s", err)
	}
	defer conn.Close()

	// Crear un canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error de canal: %s", err)
	}
	defer ch.Close()

	// Declarar la cola
	q, err := ch.QueueDeclare(
		"hello", // Nombre de la cola
		false,   // Durable
		false,   // AutoDelete
		false,   // Exclusive
		false,   // NoWait
		nil,     // Argumentos
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola: %s", err)
	}

	// Enviar un mensaje
	body := "Hello from Go Producer!"
	err = ch.Publish(
		"",     // Exchange
		q.Name, // RoutingKey
		false,  // Mandatory
		false,  // Immediate
		amqp091.Publishing{ // Usar amqp091.Publishing correctamente
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Error al publicar el mensaje: %s", err)
	}
	fmt.Printf("[x] Enviado: %s\n", body)
}

