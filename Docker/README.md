# Usando compose para desplegar los contenedores

Para desplegar usando compose.

Antes de ejecutar el compose es necesario construir las imagenes de los consumidores (recieve, subscribe) y publicadores (publish, send)

```bash
docker compose -f docker-compose-db.yml -p pubsub up

```