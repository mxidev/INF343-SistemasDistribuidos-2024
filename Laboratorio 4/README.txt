Félix Bastías       201904558-k
Maximiliano Tapia   202073552-2

Para ejecutar se debe estar en una terminal en la dirección de la carpeta tierra (dentro de grpc-server) y ocupar los comandos:
go mod init grpc
docker build . -t test-main:latest
docker run -i -p 3000:8080 test-main:latest

go run equipos.go