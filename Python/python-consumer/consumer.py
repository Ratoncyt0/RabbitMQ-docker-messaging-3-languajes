import pika

# Establecer la conexión con RabbitMQ
connection = pika.BlockingConnection(pika.ConnectionParameters('rabbitmq'))  # Nombre del contenedor RabbitMQ
channel = connection.channel()

# Declarar la cola desde la que se va a consumir los mensajes
channel.queue_declare(queue='hello')

# Método de callback para manejar los mensajes recibidos
def callback(ch, method, properties, body):
    print(f" [x] Received {body.decode()}")

# Consumir los mensajes de la cola
channel.basic_consume(queue='hello', on_message_callback=callback, auto_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
