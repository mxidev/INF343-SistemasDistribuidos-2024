# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY server.go .



RUN go build -o server .

EXPOSE 3000

CMD ["/app/server"]
