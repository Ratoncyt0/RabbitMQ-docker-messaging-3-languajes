# Como construir la imagen y enviar mensajes

Ejecutar en el directorio actual
## Construccion de la imagen
```bash
docker build -t node_send .
```

## ¿Como envio un mensaje?

```bash
docker run -e IP="172.17.0.1" -e MSG="Hola Mundo desde js" --rm node_send
```
Reemplaza la ip 172.17.0.1 (Localhost para entornos docker) por la ip de la maquina que este ejecutando RabbitMQ, y el hola mundo por el mensaje deseado. La etiqueta --rm se encargara de desaparecer al contenedor una vez enviado el mensaje