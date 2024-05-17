# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY . .



RUN go build -o datanode1 .

EXPOSE 3091

CMD ["/app/datanode1"]
