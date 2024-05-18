package main

import (
	"bufio"
	"fmt"
	pb "grpc/proto"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

// Struct para el modelo del servidor Director
type DirectorServer struct {
	pb.UnimplementedMessageServiceServer
}

func (s *DirectorServer) RequestInformation(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	if req.Body == "OK" { // Si recibo un "OK", entonces manda "NICE"
		return &pb.Message{Body: "NICE"}, nil
	}
	return &pb.Message{Body: "STOP"}, nil
}

func (s *DirectorServer) RequestMount(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	// Retorna el amount entregado por DoshBank
	amount := queryDoshBank()
	return &pb.Message{Body: amount}, nil
}

func queryDoshBank() string {
	// Implementacion para la consulta al DoshBank
	conn, err := grpc.Dial("doshbankHost:port", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error conectando con DoshBank: %v\n", err)
	}
	defer conn.Close()

	// Referencia al servidor del DoshBank para enviar la consulta
	c := pb.NewMessageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Se envia la consulta
	response, err := c.RequestMount(ctx, &pb.Message{Body: "AMOUNT"})
	if err != nil {
		fmt.Printf("Error consultando al DoshBank: %v\n", err)
	}
	return response.Body
}

func main() {
	continuar := true
	i := 0

	// Conn para mercenario 1
	conn1, err1 := grpc.Dial("merc1:port", grpc.WithInsecure())
	if err1 != nil {
		fmt.Println("Error en la conexion con host:", err1)
	}
	defer conn1.Close()

	// Se crea la referencia de conexion para mercenario 1
	client1 := pb.NewMessageServiceClient(conn1)
	response1, err := client1.RequestInformation(context.Background(), &pb.Message{Body: "REQ"})
	if err != nil {
		fmt.Printf("Error al enviar mensaje: %v\n", err)
	}

	if response1.Body == "OK" {
		fmt.Printf("Mercenario %d está OK\n", i)
		i++
	}

	if i == 7 {
		for continuar {
			menuInicio := "\n     ====================\n    |1. Iniciar Mision   |\n    |2. Terminar programa|\n     ====================\nIngrese una opcion: "
			fmt.Println(menuInicio)

			// Se pide un dato de entrada para iniciar la mision o salir
			reader := bufio.NewReader(os.Stdin)
			opcion, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error con leer un valor por consola:", err)
			}
			opcion = strings.TrimSpace(opcion)

			switch opcion {
			case "1":
				fmt.Println("Iniciando mision ...")
				menuOpciones := "\n     ============================\n    |1. Avanzar al siguiente piso|\n    |2. Volver al menu anterior  |\n     ============================\nIngrese una opcion: "
				fmt.Println(menuOpciones)

				// Se pide otro dato de entrada para avanzar de piso o volver al menu anterior
				seguir, er := reader.ReadString('\n')
				if er != nil {
					fmt.Println("Error con leer el segundo valor por consola:", er)
				}
				seguir = strings.TrimSpace(seguir)

				switch seguir {
				case "1":
					// Proceso de pisos, mercenarios y demas
					client1.RequestInformation(context.Background(), &pb.Message{Body: "NICE"}) // Aqui tenemos que ver bien como enviar el mensaje para que retorne una decision para el piso siguiente.
					// Quizas algo como tener un contador de pisos que cuando incremente avisarle al Mercenario que ahora cambio de piso.
					fmt.Println("Avanzando al siguiente piso ...")
					// Asumiendo que ya implementamos lo de arriba
					connNN, errNN := grpc.Dial("namenodeHost:port", grpc.WithInsecure())
					if errNN != nil {
						fmt.Println("Error en la conexion con host:", errNN)
					}
					defer connNN.Close()

					clientNN := pb.NewMessageServiceClient(connNN)
					clientNN.RequestInformation(context.Background(), &pb.Message{Body: "0:mercenario,piso,decision"}) // Si es 0, entonces envia informacion. Si es 1, entonces pide informacion.

				case "2":
					fmt.Println("Volviendo al menu anterior ...")
				default:
					fmt.Println("No se ingreso una opcion valida. Se volvera al menu anterior.")
					fmt.Println()
				}

			case "2":
				fmt.Println("Terminando programa ...")
				os.Exit(1)
			default:
				fmt.Println("Opcion invalida, intentelo nuevamente.")
			}
		}
		fmt.Println("Programa finalizado!")
	}

}
