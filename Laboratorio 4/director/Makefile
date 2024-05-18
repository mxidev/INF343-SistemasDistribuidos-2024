DOCKERFILE = Dockerfile

build:
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	docker build . -t director:latest

run: 
	docker run -i -p 3000:8080 director:latest