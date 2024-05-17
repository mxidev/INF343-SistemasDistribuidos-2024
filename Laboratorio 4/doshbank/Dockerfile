# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY . .



RUN go build -o doshbank .

EXPOSE 3030

CMD ["/app/doshbank"]
