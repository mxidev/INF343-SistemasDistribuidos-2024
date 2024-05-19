# Dockerfile para el servidor central
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY . .



RUN go build -o namenode .

EXPOSE 3060

CMD ["/app/namenode"]
