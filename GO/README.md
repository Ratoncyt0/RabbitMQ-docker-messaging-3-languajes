
# PracticaRabbitmq
Este proyecto muestra cómo implementar un sistema de mensajería con RabbitMQ, Docker, Docker Compose y Go.

## Requisitos

Docker y Docker Compose instalados.

Golang instalado en tu sistema.
```Bash
PracticaRabbitmq/
│
├── docker-compose.yml
├── go.mod
├── go.sum
├── go-producer/
│   └── producer.go
├── go-consumer/
│   └── consumer.go
```

## Uso

Clonar el repositorio:

```bash
$ git clone <url-del-repositorio>
$ cd PracticaRabbitmq
```

Construir y levantar contenedores:
```bash
$ docker-compose up -d --build
```

Verificar que los servicios estén corriendo:
```bash
$ docker-compose ps
```
Ver logs del productor:
```
$ docker-compose logs -f producer
```
Ver logs del consumidor:
```
$ docker-compose logs -f consumer
```
## Comunicación entre máquinas

Si deseas que el productor y consumidor se conecten a RabbitMQ en otra máquina, debes actualizar la variable de entorno RABBITMQ_HOST en el docker-compose.yml:
```
$ docker-compose up -d --build
```
```
environment:
  RABBITMQ_HOST: 192.168.1.10  # La IP de la máquina que ejecuta RabbitMQ
```
Luego, reconstruir y levantar contenedores:
```
$ docker-compose up -d --build
```
