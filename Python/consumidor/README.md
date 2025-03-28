# Recibir mensajes en mi maquina local 

1. Primero tenemos que crear la imagen con el Dockerfile:

Para crear una imagen con el dockerfile para este apartado CONSUMIDOR:
```bash
docker build -t consumerpy .
```
2. Para recibir mensajes de cualquier maquina usando la imagen: 
```bash
docker run --rm -it consumerpy
```
Ya con esto podemos obsevar como recibimos los mensajes