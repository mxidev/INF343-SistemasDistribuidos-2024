# Dockerfile para el servidor central
FROM golang:1.22.2

WORKDIR /app

COPY . .



RUN go build -o datanode3 .

EXPOSE 3093

CMD ["/app/datanode3"]
