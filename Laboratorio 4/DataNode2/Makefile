DOCKERFILE = Dockerfile

build:
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go get grpc/proto
	docker build . -t datanode2:1.22.2

run: 
	docker run -p 3092:8080 datanode2:1.22.2