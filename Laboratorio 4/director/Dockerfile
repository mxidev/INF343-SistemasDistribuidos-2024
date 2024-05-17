# Dockerfile para el servidor central
FROM golang:latest

WORKDIR /app

COPY . .



RUN go build -o director .

EXPOSE 3000

CMD ["/app/director"]
