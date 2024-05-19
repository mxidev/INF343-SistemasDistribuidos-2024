# Dockerfile para el servidor central
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY . .


RUN go mod download && go mod verify
RUN go build -o doshbank .

EXPOSE 3030

CMD ["/app/doshbank"]
