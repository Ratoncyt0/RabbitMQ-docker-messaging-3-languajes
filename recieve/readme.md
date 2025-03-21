# Construir imagen y usar el contenedor para recibir mensajes

## Para construir la imagen
```bash
docker build -t node_recieve .
```
## ¿Como usar la imagen para recibir mensajes?
```bash
docker run -e IP="192.168.234.38" node_recieve
```