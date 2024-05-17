# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY . .



RUN go build -o datanode2 .

EXPOSE 3092

CMD ["/app/datanode2"]
