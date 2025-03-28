import pika

# Dirección IP o nombre del host de la máquina que corre RabbitMQ
rabbitmq_host = '10.6.101.98'

# Conexión al servidor RabbitMQ remoto
connection = pika.BlockingConnection(pika.ConnectionParameters(host=rabbitmq_host))
channel = connection.channel()

# Asegurarse de que la cola existe
channel.queue_declare(queue='hello')

# Función que maneja la llegada de mensajes
def callback(ch, method, properties, body):
    print(f" [x] Received {body}")

# Consumir mensajes de la cola
channel.basic_consume(queue='hello', on_message_callback=callback, auto_ack=True)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
