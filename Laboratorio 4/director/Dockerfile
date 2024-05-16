# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY . .



RUN go build -o tierra .

EXPOSE 3000

CMD ["/app/tierra"]
