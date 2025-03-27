# Usar una imagen de Go
FROM golang:1.20-alpine

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del proyecto
COPY . .

# Instalar las dependencias
RUN go mod tidy

# Compilar el código
RUN go build -o producer .

# Ejecutar el productor
CMD ["./producer"]
