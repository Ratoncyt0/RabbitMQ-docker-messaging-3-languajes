FROM node:16

# Creamos el directorio
RUN mkdir /usr/src/app

# Establecemos el directorio de trabajo
WORKDIR /usr/src/app

# Instalamos las dependencias necesarias, na mas amqplib
RUN npm install amqplib

# Copiamos los archivos dentro del contenedor
COPY package.json .
COPY send.js .

# Nos aseguramos que runsend tenga permisos de ejecucion

# Nos aseguramos que send.js se ejecute cada vez que se inicia el contenedor
CMD ["node", "send.js"]
