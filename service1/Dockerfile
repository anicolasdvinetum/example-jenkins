FROM golang:1.21-alpine

WORKDIR /app

COPY main.go .

RUN go build -o service1

EXPOSE 8080

CMD ["./service1"]