# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY . .



RUN go build -o namenode .

EXPOSE 3060

CMD ["/app/namenode"]
