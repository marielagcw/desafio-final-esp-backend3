# Primer paso: construir la aplicación
FROM golang:latest AS build

# Establecer el directorio de trabajo
WORKDIR /go/src/app

# Copiar el código fuente a la imagen
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# Segundo paso: ejecutar la aplicación compilada
FROM alpine:latest

# Copiar la aplicación compilada del primer paso
COPY --from=build /go/src/app/app /app
COPY .env .env

# Establecer el directorio de trabajo
WORKDIR /

# Ejecutar la aplicación
CMD ["/app"]