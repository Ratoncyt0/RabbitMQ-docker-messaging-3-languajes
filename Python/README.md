# RabbitMQ-docker-messaging-Python

# Crea y ejecuta RabbitMQ con doker

```bash
docker pull rabbitmq:management
docker run -d --name rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:management
```


# Luego Creamos la carpeta python-producer y python-consumer.
```bash
mkdir python-producer
```
Para el productor 
```bash
mkdir python-consumer
```