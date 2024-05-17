package proto

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
var DoshBankSync MessageServiceClient

func init() {
	// Se crean las conexiones con cada "cliente" que interactua
	// de forma sincrona con El Director
	NameNodeClientConn, err1 := grpc.Dial("localhost:50051", grpc.WithInsecure())
	MercenaryClientConn, err2 := grpc.Dial("localhost:50052", grpc.WithInsecure())
	DoshBankSyncConn, err3 := grpc.Dial("localhost:50053", grpc.WithInsecure())

	if err1 != nil {
		log.Fatalf("No se pudo conectar con el NameNode: %s\n", err1)
	}
	if err2 != nil {
		log.Fatalf("No se pudo conectar con los Mercenarios: %s\n", err2)
	}
	if err3 != nil {
		log.Fatalf("No se pudo conectar con el DoshBank: %s\n", err3)
	}

	NameNodeClient = NewMessageServiceClient(NameNodeClientConn)
	MercenaryClient = NewMessageServiceClient(MercenaryClientConn)
	DoshBankSync = NewMessageServiceClient(DoshBankSyncConn)
}

// Override a los metodos de la interfaz
func (s *Server) RequestMount(ctx context.Context, message *Message) (*Message, error) {
	fmt.Printf("Solicitud recibida por El Director: %s\n", message.Body)
	return message, nil // Por mientras
}

func (s *Server) RequestDecisionToMercenary(ctx context.Context, message *Message) (*Message, error) {
	fmt.Printf("Solicitud de Decision recibida: %s\n", message.Body)
	return &Message{Body: "OK"}, nil // Por mientras
}

func (s *Server) RequestInformationToDataNode(ctx context.Context, message *Message) (*Message, error) {
	fmt.Printf("Solicitud para enviar informacion de NameNode recibida: %s\n", message.Body)

	fileName := strings.Split(message.Body, ",") // Asumiendo formato merc,piso,decs
	mercenary := fileName[0]
	floor := fileName[1]

	// Se abre el archivo con el nombre correspondiente
	file, err := os.Open(mercenary + "_" + floor + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Creamos un scanner para leer lineas del archivo
	scanner := bufio.NewScanner(file)
	response := ""

	// Se itera por cada linea
	for scanner.Scan() {
		linea := scanner.Text()
		segment := strings.Split(linea, " ") // Formato: * num
		response += segment[1] + ","         // Interesa la decision que hizo
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return &Message{Body: mercenary + "," + floor + ":" + response}, nil // Por mientras
	// Se retornaria algo como "JOHN,3:7,4,1,10,4,"
}

func (s *Server) AddInformation(ctx context.Context, message *Message) (*empty.Empty, error) {
	fmt.Println("Instruccion de agregar informacion de NameNode recibida")
	body := message.Body
	mercenary := strings.Split(body, ",") // Asumiendo formato merc,piso,decs
	fileName := mercenary[0] + "_" + mercenary[1] + ".txt"

	// Se abre el archivo con el formato de nombre
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error con abrir el archivo: ", err)
	}
	defer file.Close()

	// Se escriben datos en el archivo
	_, err = fmt.Fprintf(file, "* %s\n", mercenary[2])
	if err != nil {
		fmt.Println("Error al escribir en el archivo: ", err)
	}

	fmt.Println("Informacion almacenada, mensaje enviado.")
	return &emptypb.Empty{}, nil
}
