Por ultimo crea una carpeta python-consumer la cual sera el consumidor final: 

1.  Crea un archivo consumer.py en en esta carpeta con su respectivo codigo.
2.  Crea un archivo Dockerfile en el directorio del consumidor Python con su respectivo codigo.

Por ultimo en el directorio del consumidor, ejecuta estos comandos para construir y ejecutar los contenedores:


```bash
docker build -t python-consumer .
docker run --link rabbitmq:rabbitmq python-consumer
```

