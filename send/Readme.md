# Como construir la imagen y enviar mensajes

Ejecutar en el directorio actual
## Construccion de la imagen
```bash
docker build -t node_send .
```

## ¿Como envio un mensaje?

```bash
docker run -d -e IP="192.168.1.1" -e MSG="Hola Mundo" --rm node_send
```
Reemplaza la ip 192.168.1.1 por la ip de la maquina que este ejecutando RabbitMQ, y el hola mundo por el mensaje deseado. La etiqueta --rm se encargara de desaparecer al contenedor una vez enviado el mensaje