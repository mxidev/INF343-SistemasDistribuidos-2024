package proto

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	UnimplementedMessageServiceServer
}

var NameNodeClient MessageServiceClient
var MercenaryClient MessageServiceClient
var DoshBankAsync MessageServiceClient

func init() {
	NameNodeClientConn, err1 := grpc.Dial("localhost:50051", grpc.WithInsecure())
	MercenaryClientConn, err2 := grpc.Dial("localhost:50052", grpc.WithInsecure())
	DoshBankAsyncConn, err3 := grpc.Dial("localhost:50053", grpc.WithInsecure())

	if err1 != nil {
		fmt.Printf("No se pudo conectar con el NameNode: %s\n", err1)
	}
	if err2 != nil {
		fmt.Printf("No se pudo conectar con los Mercenarios: %s\n", err2)
	}
	if err3 != nil {
		fmt.Printf("No se pudo conectar con el DoshBank: %s\n", err3)
	}

	NameNodeClient = NewMessageServiceClient(NameNodeClientConn)
	MercenaryClient = NewMessageServiceClient(MercenaryClientConn)
	DoshBankAsync = NewMessageServiceClient(DoshBankAsyncConn)
}

// Override a los metodos de la interfaz
func (s *Server) RequestDoshBank(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Solicitud recibida por El Director: %s", message.Body)
	return message, nil // Por mientras
}

func (s *Server) RequestDirector(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Solicitud recibida por Mercenario: %s", message.Body)
	return message, nil // Por mientras
}

func (s *Server) RequestDecision(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Decision recibida por Mercenario: %s", message.Body)
	return message, nil // Por mientras
}

func (s *Server) AddInformation(ctx context.Context, message *Message) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}
