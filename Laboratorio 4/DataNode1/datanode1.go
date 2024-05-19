package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

type DN1Server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *DN1Server) RequestInformation(ctx context.Context, message *pb.Message) (*pb.Message, error) {

	// Al DataNode le llega un mensaje con formato "id:m,p,d"
	// 1 es para guardar la informacion en un archivo
	// 0 para retornar la informacion del archivo
	split := strings.Split(message.Body, ":")
	id := split[0]
	slice := split[1]
	body := strings.Split(slice, ",")
	msg := ""

	mercenary := body[0]
	floor := body[1]
	decision := body[2]

	fileName := mercenary + "_" + floor + ".txt"

	if id == "0" {
		// Retornar datos segun el nombre de mercenario y piso
		msg = transferData(mercenary, floor)
		fmt.Println("Solicitud de NameNode recibida. Operacion completada:", msg)
	}

	if id == "1" {
		// Se envian los datos para guardarlos en algun Data Node
		// Se abre el archivo con el formato de nombre
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error con abrir el archivo: ", err)
		}
		defer file.Close()

		// Se escribe en el archivo
		file.WriteString("* " + decision + "\n")
		fmt.Println("Se escribio la data: *", decision)
	}
	return &pb.Message{Body: msg}, nil
}

func transferData(name string, floor string) string {

	fName := name + "_" + floor + ".txt"
	fileOpen, err := os.Open(fName)
	if err != nil {
		fmt.Println("Error abriendo el archivo: ", err)
	}
	defer fileOpen.Close()

	retorno := ""
	scann := bufio.NewScanner(fileOpen)
	for scann.Scan() {
		split := strings.Split(scann.Text(), " ")
		retorno += split[1] + "," // Formato: d1,d2,...,dn,
	}

	return retorno
}

func main() {

	list, err := net.Listen("tcp", ":3060")
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()
	for {
		pb.RegisterMessageServiceServer(serv, &DN1Server{})
		if err = serv.Serve(list); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
	}
}
