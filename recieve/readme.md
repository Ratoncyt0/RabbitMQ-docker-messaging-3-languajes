# Construir imagen y usar el contenedor para recibir mensajes

## Para construir la imagen
```bash
docker build -t node_recieve .
```
## ¿Como usar la imagen para recibir mensajes?
```bash
docker run -e IP="192.168.1.1" node_recieve
```
Cambia la ip 192.168.1.1 por la ip de la maquina que este ejecutando RabbitMQ, se abrira la consola de node ejecutando recieve.js