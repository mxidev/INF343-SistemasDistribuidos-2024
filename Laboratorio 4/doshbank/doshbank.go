package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

type DoshBankServer struct {
	pb.UnimplementedMessageServiceServer
	totalAmount int
}

func (s *DoshBankServer) RequestMount(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	if req.Body == "AMOUNT" { // Si recibe un "AMOUNT", retorna el monto
		return &pb.Message{Body: fmt.Sprintf("%d", s.totalAmount)}, nil
	}
	return &pb.Message{Body: "Invalid Request"}, nil
}

func (s *DoshBankServer) AddAmount(amount int) {
	// Funcion auxiliar para incrementar el monto acumulado
	s.totalAmount += amount
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		fmt.Printf("Error esperando respuesta: %v\n", err)
	}

	// Se crea la referencia para el servidor del DoshBank
	s := grpc.NewServer()
	doshBankServer := &DoshBankServer{}
	pb.RegisterMessageServiceServer(s, doshBankServer)
	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("Error en el servicio: %v\n", err)
		}
	}()

	doshBankServer.AddAmount(100) // De prueba
}
