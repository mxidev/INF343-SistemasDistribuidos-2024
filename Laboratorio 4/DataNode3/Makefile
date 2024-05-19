DOCKERFILE = Dockerfile

build:
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go get grpc/proto
	docker build . -t datanode3:1.22.2

run: 
	docker run -p 3093:8080 datanode3:1.22.2