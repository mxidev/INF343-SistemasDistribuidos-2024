package main

import (
	"bufio"
	"fmt"
	pb "grpc/proto"
	"math/rand"
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
	if req.Body == "AMOUNT" {
		// Llamar al Dosh Bank
		query := queryDoshBank()
		return &pb.Message{Body: query}, nil
	} else {
		// Formato m:p:d
		fmt.Println("MPD")
	}
	return &pb.Message{Body: "STOP"}, nil
}

func queryDoshBank() string {
	// Implementacion para la consulta al DoshBank
	conn, err := grpc.Dial("dist077:3030", grpc.WithInsecure())
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
	conn1, err1 := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err1 != nil {
		fmt.Println("Error en la conexion con host:", err1)
	}
	defer conn1.Close()

	// Se crea la referencia de conexion para mercenario 1
	client1 := pb.NewMessageServiceClient(conn1)
	response1, err := client1.RequestInformation(context.Background(), &pb.Message{Body: "REQ"})
	if err != nil {
		panic("Error al enviar mensaje al jugador 1")
	}

	if response1.Body == "OK" {
		fmt.Printf("Mercenario %d está OK\n", i)
		i++
	}

	if i == 1 {
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
			case "1": //Piso1
				fmt.Println("Iniciando mision ...")
				rand.Seed(time.Now().UnixNano())
				//llega primera decision
				var dec string
				x := rand.Intn(100)
				if dec == "A" && x < 35 && x >= 0 {
					fmt.Print("VIVO")
				} else if dec == "B" && x < 70 && x >= 35 {
					fmt.Print("VIVO")
				} else if dec == "C" && x <= 100 && x >= 70 {
					fmt.Print("VIVO")
				} else {
					fmt.Print("MUERTO")
				}

				menuOpciones := "\n     ============================\n    |1. Avanzar al siguiente piso|\n    |2. Consultar Decisiones  |\n     ============================\nIngrese una opcion: "
				fmt.Println(menuOpciones)

				// Se pide otro dato de entrada para avanzar de piso o volver al menu anterior
				seguir, er := reader.ReadString('\n')
				if er != nil {
					fmt.Println("Error con leer el segundo valor por consola:", er)
				}
				seguir = strings.TrimSpace(seguir)

				switch seguir { //Piso2
				case "1":
					// Proceso de pisos, mercenarios y demas
					client1.RequestInformation(context.Background(), &pb.Message{Body: "GO"}) // Aqui tenemos que ver bien como enviar el mensaje para que retorne una decision para el piso siguiente.
					// Quizas algo como tener un contador de pisos que cuando incremente avisarle al Mercenario que ahora cambio de piso.
					fmt.Println("Avanzando al siguiente piso ...")
					// Asumiendo que ya implementamos lo de arriba
					tYt := rand.Intn(2)
					if tYt == 1 {
						fmt.Println("El pasillo A es el correcto..")
					} else {
						fmt.Println("El pasillo B es el correcto..")
					}
					var decp2 int
					if decp2 == 1 && tYt == 1 {
						fmt.Print("VIVO")
					} else if decp2 == 2 && tYt == 2 {
						fmt.Print("VIVO")
					} else {
						fmt.Print("MUERTO")
					}

					connNN, errNN := grpc.Dial("namenodeHost:3060", grpc.WithInsecure())
					if errNN != nil {
						fmt.Println("Error en la conexion con host:", errNN)
					}
					defer connNN.Close()

					clientNN := pb.NewMessageServiceClient(connNN)
					clientNN.RequestInformation(context.Background(), &pb.Message{Body: "0:mercenario,piso,decision"}) // Si es 0, entonces envia informacion.
				case "2":
					connNN, errNN := grpc.Dial("dist077:3060", grpc.WithInsecure())
					if errNN != nil {
						fmt.Println("Error en la conexion con host:", errNN)
					}
					defer connNN.Close()

					clientNN := pb.NewMessageServiceClient(connNN)
					clientNN.RequestInformation(context.Background(), &pb.Message{Body: "1:mercenario,piso,decision"}) // Si es 1, entonces pide informacion.
				default:
					fmt.Println("Volviendo al menu anterior ...")
					fmt.Println()
				}

			case "2":
				fmt.Println("Terminando programa ...")
				os.Exit(1)
			default:
				fmt.Println("Opción invalida, intentelo nuevamente.")
			}
		}
		fmt.Println("Programa finalizado!")
	}

}
