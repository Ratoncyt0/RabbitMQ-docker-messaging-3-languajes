# Integrantes
* Camilo Andres Sanmiguel Uribe - 2200922
* Brayan Yecid Aparicio Goyeneche - 2205089
* Zamir Francisco Granados Peñaloza - 2211888


### RabbitMQ docker messaging en tres lenguajes (js, py, GO)
An example of hello world using docker
based on https://www.rabbitmq.com/tutorials/

# Despliegue sin compose
## Primero iniciamos el contenedor de RabbitMQ
```bash
docker run -d -it --rm --name rabbitmq -p 5552:5552 -p 15672:15672 -p 5672:5672 -e RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS='-rabbitmq_stream advertised_host localhost' rabbitmq:4-management
```
## Recomendado (Activar los servicios de streamming)
```bash
docker exec rabbitmq rabbitmq-plugins enable rabbitmq_stream rabbitmq_stream_management 
```
Dentro de los productores (publish o send) y consumidores (recieve o subscribe) correspondientes a cada lenguaje se encuentran las instrucciones de construccion de la imagen y despliege de contenedor sin uso de compose.

# Despliegue usando compose
Para mas informacion dirijase al directorio Docker
