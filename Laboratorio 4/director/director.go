package main

import (
	"bufio"
	"fmt"
	pb "grpc/proto"
	"os"
	"strings"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func main() {
	continuar := true
	i := 0

	// Conn para mercenario 1
	conn1, err1 := grpc.Dial("host:port", grpc.WithInsecure())
	if err1 != nil {
		fmt.Println("Error en la conexion con host:", err1)
	} else {
		j := pb.NewMessageServiceClient(conn1)
		response1, _ := j.RequestDecisionToMercenary(context.Background(), &pb.Message{
			Body: "REQ",
		})
		if response1.Body == "OK" {
			i++
		}
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
					j := pb.NewMessageServiceClient(conn)
					j.RequestDecisionToMercenary()
					fmt.Println("Avanzando al siguiente piso ...")
					// Proceso de pisos, mercenarios y demas
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
