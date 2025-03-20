FROM node:16

# Creamos el directorio
RUN mkdir /usr/src/app

# Establecemos el directorio de trabajo
WORKDIR /usr/src/app

# Instalamos las dependencias necesarias, na mas amqplib
RUN npm install amqplib

# Copiamos los archivos dentro del contenedor
COPY . .

# Aseguramos que el archivo runsend (sh) tenga permisos de ejecución
RUN chmod +x ./runsend

CMD [ "/usr/src/app/runsend" ]