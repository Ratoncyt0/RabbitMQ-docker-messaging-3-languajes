# Usando compose para desplegar los contenedores

Para desplegar usando compose.

Antes de ejecutar el compose es necesario construir las imagenes de los consumidores (recieve, subscribe) y publicadores (publish, send)

## Version 1 LOCAL
```bash
docker compose -f docker-compose.yml -p pubsub up

```

## Version 2 network
Es necesario modificar el archivo "GO/go-rabbitmq/.env"
```bash
docker compose -f docker-compose.yml -p pubsub up

```