import pika

# Dirección IP o nombre del host de la máquina que corre RabbitMQ
rabbitmq_host = '10.6.101.98'  #Modificar la IP a la maquina que quiera mandar mensajes

# Conexión al servidor RabbitMQ remoto
connection = pika.BlockingConnection(pika.ConnectionParameters(host=rabbitmq_host))
channel = connection.channel()

# Asegurarse de que la cola existe
channel.queue_declare(queue='hello')

# Enviar un mensaje a la cola
channel.basic_publish(exchange='',
                      routing_key='hello',
                      body='Hola desde Python de Zamir - SIU!')

print(" [x] Sent 'Hola desde Python de Zamir - SIU!'")

# Cerrar la conexión
connection.close()