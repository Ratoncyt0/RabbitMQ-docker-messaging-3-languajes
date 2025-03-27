import pika

# Establecer la conexión con RabbitMQ
connection = pika.BlockingConnection(pika.ConnectionParameters('rabbitmq'))  # Nombre del contenedor RabbitMQ
channel = connection.channel()

# Crear una cola llamada 'hello'
channel.queue_declare(queue='hello')

# Enviar un mensaje a la cola
channel.basic_publish(exchange='',
                      routing_key='hello',
                      body='Hello from Python Producer!')

print(" [x] Sent 'Hello from Python Producer!'")

# Cerrar la conexión
connection.close()
