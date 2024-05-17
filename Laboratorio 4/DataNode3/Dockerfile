# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY . .



RUN go build -o datanode3 .

EXPOSE 3093

CMD ["/app/datanode3"]
