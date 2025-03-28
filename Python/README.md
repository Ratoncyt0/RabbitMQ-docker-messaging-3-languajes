# RabbitMQ-docker-messaging-Python

# Ejecuta RabbitMQ con doker

En el directorio de tu elección, crea y ejecuta:

```bash
docker pull rabbitmq:management
docker run -d --name rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:management
```

# Instalamos las dependencias de Python
```bash
pip install pika
```

# Creamos una carpeta "Consumidor y Productor"
productor:
```bash
mkdir productor
``` 
consumidor:
```bash
mkdir consumidor
```
# Carpeta Productor 
1. Creamos un --> nano Dockerfile
2. Creamos un --> nano send.py

# Carpeta Consumidor 
1. Creamos un --> nano Dockerfile
2. Creamos un --> nano receive.py
