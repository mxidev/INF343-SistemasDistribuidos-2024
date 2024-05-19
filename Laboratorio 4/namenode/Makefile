DOCKERFILE = Dockerfile

build:
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go get grpc/proto
	docker build . -t namenode:1.22.2

run: 
	docker run -p 3060:8080 namenode:1.22.2