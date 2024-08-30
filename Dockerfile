# Etapa 1: Compilación de la aplicación
FROM golang:1.23 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de Go mod y sum y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente de la aplicación
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Etapa 2: Creación de la imagen final
FROM alpine:latest

# Establecer variables de entorno para asegurar un entorno de ejecución seguro
ENV GO_ENV=production

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copiar la aplicación compilada desde la etapa de compilación
COPY --from=builder /app/app .

# Exponer el puerto en el que la aplicación va a correr (ajustar según sea necesario)
EXPOSE 8080

# Comando de inicio de la aplicación
CMD ["./app"]
