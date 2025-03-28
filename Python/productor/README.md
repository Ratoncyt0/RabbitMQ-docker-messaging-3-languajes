# Enviar mensajes a otras maquinas 

1. Primero tenemos que crear la imagen con el Dockerfile:

Para crear una imagen con el dockerfile para este apartado PRODUCTOR:
```bash
docker build -t sendpy .
```
2. Para enviar mensajes a la maquina usando la imagen: 
```bash
python send.py
```
Ya con esto enviamos el mensaje y en consola nos sale el mensaje que enviamos. 

# Modificar mensaje y IP de la maquina a enviar

1. Entramos al nano "send.py" y modificamos la ip a la que queremos mandar el mensaje.
```bash
rabbitmq_host = '10.6.101.91'  #Modificar la IP a la maquina que quiera mandar mensajes
```
2. Alli mismo en el nano "send.py" podemos modificar el mensaje en el body:
```bash
body='Hola desde Python de Zamir - SIU!'
```