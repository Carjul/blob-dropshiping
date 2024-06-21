# Establecer la imagen base
FROM golang:1.22.4

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum al directorio de trabajo
COPY go.mod go.sum ./

# Descargar las dependencias del módulo
RUN go mod download

# Copiar el resto de los archivos al directorio de trabajo
COPY . .

# Compilar la aplicación
RUN go build -o app ./cmd/main.go

# Exponer el puerto en el que se ejecuta la aplicación
EXPOSE 8080

# Ejecutar la aplicación cuando se inicie el contenedor
CMD ["./app"]