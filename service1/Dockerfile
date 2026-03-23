FROM golang:1.21-alpine

WORKDIR /app

# Copiar go.mod y go.sum primero (si existe), para cachear dependencias.
COPY go.mod ./

RUN go mod download

# Ahora sí copiamos el código
COPY . .

RUN go build -o service1

EXPOSE 8080

CMD ["./service1"]