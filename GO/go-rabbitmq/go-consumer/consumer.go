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

	// Consumir el mensaje
	msgs, err := ch.Consume(
		q.Name, // Nombre de la cola
		"",     // Consumer tag
		true,   // AutoAck
		false,  // Exclusive
		false,  // NoLocal
		false,  // NoWait
		nil,    // Argumentos
	)
	if err != nil {
		log.Fatalf("Error al consumir el mensaje: %s", err)
	}

	// Procesar los mensajes recibidos
	for msg := range msgs {
		fmt.Printf("[x] Recibido: %s\n", msg.Body)
	}
}
