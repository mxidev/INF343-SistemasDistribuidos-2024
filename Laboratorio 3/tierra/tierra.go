package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "grpc-server/munition"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

// Se define la estructura del servidor.
// Con esto podemos acceder a los servicios y mensajes.
// (yo lo entiendo como los metodos y atributos de un objeto)
type Server struct {
	pb.UnimplementedMunitionServiceServer
}

var mp int = 0
var at int = 0

func solicitarM(id int32, solAT int, solMP int) bool {
	if solAT < at && solMP < mp {
		at = at - solAT
		mp = mp - solMP
		return true
	} else {
		return false
	}
}

func periodicFunction() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:

			if at+10 >= 50 {
				at = 50
			} else {
				at = at + 10
			}

			if mp+5 >= 20 {
				mp = 20
			} else {
				mp = mp + 5
			}
			fmt.Println("Función activada", "MP: ", mp, " AT: ", at)
			fmt.Println("")
		}
	}
}

// Este es el @Override del servicio descrito en la interfaz del protobuf (munition.proto)
// El return debe seguir el formato del mensaje Munition Response
func (s *Server) RequestMunition(ctx context.Context, in *pb.MunitionRequest) (*pb.MunitionResponse, error) {
	// Se guarda la respuesta de solicitarM
	solicitud := solicitarM(int32(in.IdGrupo), int(in.CantidadAt), int(in.CantidadMp))
	if solicitud {
		// Caso positivo
		fmt.Printf("Recepción de solicitud desde equipo %d, %d AT y %d MP --ACEPTADA--\n", int(in.IdGrupo)+1, int(in.CantidadAt), int(in.CantidadMp))
		fmt.Printf("AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", at, mp)
		return &pb.MunitionResponse{Response: true, DisponibleAt: int32(at), DisponibleMp: int32(mp)}, nil
	} else {
		// Caso negativo
		fmt.Printf("Recepción de solicitud desde equipo %d, %d AT y %d MP --DENEGADA--\n", int(in.IdGrupo)+1, int(in.CantidadAt), int(in.CantidadMp))
		fmt.Printf("AT EN SISTEMA: %d ; MP EN SISTEMA: %d\n", at, mp)
		return &pb.MunitionResponse{Response: false, DisponibleAt: int32(at), DisponibleMp: int32(mp)}, nil
	}
}

func main() {

	conn, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("No se pudo establecer la conexion: %s\n", err)
	}
	defer conn.Close()
	fmt.Println("Servidor corriendo! Esperando una solicitud . . .")
	go periodicFunction()
	// Se crea y registra el servidor grpc
	s := grpc.NewServer()
	pb.RegisterMunitionServiceServer(s, &Server{})
	if err := s.Serve(conn); err != nil {
		fmt.Printf("No se pudo levantar el servidor gRPC: %s", err)
	}

	time.Sleep(60 * time.Second)
}
