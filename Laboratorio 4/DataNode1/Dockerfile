# Dockerfile para el servidor central
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY . .


RUN go mod download && go mod verify
RUN go build -o datanode1 .

EXPOSE 3091

CMD ["/app/datanode1"]
