DOCKERFILE = Dockerfile

build:
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go get grpc/proto
	docker build . -t doshbank:1.22.2

run: 
	docker run -i -p 3030:8080 doshbank:1.22.2