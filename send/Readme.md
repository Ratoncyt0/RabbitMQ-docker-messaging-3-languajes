Como construir la imagen y utilizar el contenedor

docker run -e IP="192.168.234.38" -e MSG="Hola desde js" --rm --name send_node -it --entrypoint /bin/bash send_node