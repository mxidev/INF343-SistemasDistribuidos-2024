DOCKERFILE = Dockerfile

build:
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go get grpc/proto
	docker build . -t director:1.22.2

run: 
	docker run -i -p 3000:8080 director:1.22.2