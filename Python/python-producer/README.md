
Crea una carpeta python-producer en tu directorio y alli agrega:

1. Crea un archivo producer.py y agrega su respectivo codigo.

2. En el mismo directorio, crea un archivo Dockerfile para el productor Python con su respectivo codigo.


En el directorio donde tienes el Dockerfile, ejecuta estos comandos para construir y ejecutar los contenedores:

```bash
docker build -t python-producer .
docker run --link rabbitmq:rabbitmq python-producer
```