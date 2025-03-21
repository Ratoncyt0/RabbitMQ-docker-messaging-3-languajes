# RabbitMQ docker messaging in js
An example of hello world using docker
based on https://www.rabbitmq.com/tutorials/tutorial-one-javascript-stream

Primero iniciamos el contenedor
```bash
docker run -it --rm --name rabbitmq -p 5552:5552 -p 15672:15672 -p 5672:5672 -e RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS='-rabbitmq_stream advertised_host localhost' rabbitmq:4-management
```
Recomendado (Activar los servicios de streamming)
```bash
docker exec rabbitmq rabbitmq-plugins enable rabbitmq_stream rabbitmq_stream_management 
```