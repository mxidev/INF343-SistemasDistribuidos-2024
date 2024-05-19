package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

type NameNodeServer struct {
	pb.UnimplementedMessageServiceServer
	dataNodes []string
}

func (s *NameNodeServer) RequestInformation(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	// Al NameNode le llega un mensaje con formato "1:m,p,d"
	// 1 es para guardar la informacion hacia los datanodes
	// 0 para pedir la informacion de los datanodes

	// Parsear la información
	parts := strings.Split(req.Body, ":")
	id := parts[0]
	slice := strings.Split(parts[1], ",")
	mercenary := slice[0]
	piso := slice[1]
	decision := slice[2]

	// Seleccionar un DataNode al azar
	rand.Seed(time.Now().UnixNano())
	randomNode := s.dataNodes[rand.Intn(len(s.dataNodes))]

	// Crear el archivo en el NameNode
	fileName := "decisiones.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%s %s %s\n", mercenary, piso, randomNode)); err != nil {
		return nil, err
	}

	// Enviar la información al DataNode seleccionado
	conn, err := grpc.Dial(randomNode, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)
	_, err = client.RequestInformation(ctx, &pb.Message{Body: fmt.Sprintf("%s:%s,%s,%s", id, mercenary, piso, decision)})
	if err != nil {
		return nil, err
	}

	return &pb.Message{Body: "OK"}, nil
}

func main() {
	// Se escucha en un puerto de forma sincrona
	lis, err := net.Listen("tcp", ":3092")
	if err != nil {
		fmt.Printf("Error esperando respuesta: %v\n", err)
	}

	dataNodes := []string{"datanode1:3091", "datanode2:3092", "datanode3:3093"}
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &NameNodeServer{dataNodes: dataNodes})
	fmt.Printf("NameNode escuchando: %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Error en el servicio: %v\n", err)
	}
}
