# RabbitMQ-docker-messaging-Python

# Crea y ejecuta RabbitMQ con doker
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

