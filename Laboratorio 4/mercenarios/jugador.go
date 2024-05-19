package main

import (
	"fmt"
	pb "grpc/proto"
	"net"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type MercenaryServer struct {
	pb.UnimplementedMessageServiceServer
}

func sendMessage(flag string) {
	// Enviar mensaje (flag) al Director
	conn, err := grpc.Dial("directorHost:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se pudo conectar con el servidor grpc:", err)
	}
	defer conn.Close()

	c := pb.NewMessageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RequestInformation(ctx, &pb.Message{Body: flag})
	if err != nil {
		fmt.Println("No se pudo enviar la respuesta: %v\n", err)
	}
	fmt.Printf("Respuesta del servidor: %s\n", r.Body)
}

func sendQueryAmount() {
	// Obtener monto acumulado desde el Director
	conn, err := grpc.Dial("directorHost:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error al conectar con el Director: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)
	response, err := client.RequestInformation(context.Background(), &pb.Message{Body: "AMOUNT"})
	if err != nil {
		fmt.Printf("Error al consultar el monto acumulado: %v\n", err)
	}
	fmt.Printf("Monto acumulado: %s\n", response.Body)
}

func main() {
	var i int
	begin := true

	conn, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error para esperar una respuesta: ", err)
	}
	defer conn.Close()

	serv := grpc.NewServer()
	pb.RegisterMessageServiceServer(serv, &MercenaryServer{})
	if err = serv.Serve(conn); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}

	for begin {
		fmt.Print("Si está listo ingrese [1]\n")
		fmt.Scanf("%d", &i)
		if i == 1 {
			sendMessage("OK") // Se crea la conexion grpc
			begin = false
		}
	}
	flag := true
	// estado := true
	contPiso := 0
	for flag {
		var eleccion string

		// //Caso en el que el jugador haya muerto
		// if estado == false {
		// 	fmt.Print("Usted ha muerto...")
		// }

		//Opciones para el jugador
		MenuInicio := "Jugador, ¿Qué desea hacer?\n [1] Explorar piso\n [2] Ver monto acumulado\nIngrese opción:"
		fmt.Print(MenuInicio)
		fmt.Scanln(&eleccion)

		switch eleccion {
		case "1":
			var eleccionPiso string
			var eleccionPiso3 string

			//Piso 1
			if contPiso == 0 {
				flagPiso1 := true
				for flagPiso1 {
					fmt.Print("Jugador, ¿Qué desea ocupar?\n [1] Escopeta\n [2] Rifle automático\n [3] Puños eléctricos ")
					fmt.Scanln(&eleccionPiso)

					sendMessage(fmt.Sprintf("%s:%s:%s", "JG", "1", eleccionPiso))
					contPiso++
					flagPiso1 = false

					// switch eleccionPiso {
					// case "1":
					// 	//comunicar 1
					// 	sendMessage(fmt.Sprintf("%s:%s:%s", "M1", "1", eleccionPiso))
					// 	contPiso++
					// 	flagPiso1 = false

					// case "2":
					// 	//comunicar 2
					// 	sendMessage(fmt.Sprintf("%s:%s:%s", "M1", "1", eleccionPiso))
					// 	contPiso++
					// 	flagPiso1 = false
					// case "3":
					// 	//comunicar 3
					// 	sendMessage("3")
					// 	contPiso++
					// 	flagPiso1 = false
					// default:
					// 	fmt.Println("Opción no válida")
					// }
				}

				//Piso 2
			} else if contPiso == 1 {
				flagPiso2 := true
				for flagPiso2 {
					fmt.Print("Jugador, ¿Dónde desea ir?\n [1] A\n [2] B ")
					fmt.Scanln(&eleccionPiso)

					sendMessage(fmt.Sprintf("%s:%s:%s", "JG", "2", eleccionPiso))
					contPiso++
					flagPiso2 = false

					// switch eleccionPiso {
					// case "1":
					// 	//comunicar A
					// 	sendMessage("A")
					// 	contPiso++
					// 	flagPiso2 = false

					// case "2":
					// 	//comunicar B
					// 	sendMessage("B")
					// 	contPiso++
					// 	flagPiso2 = false
					// default:
					// 	fmt.Println("Opción no válida")
					// }
				}

				//Piso 3
			} else if contPiso == 2 {
				i := 0
				fmt.Print("A continuación debe elegir 5 números del 1 al 15\n")
				for i < 5 {
					var aux string
					fmt.Printf("Elija el número %d: ", i+1)
					fmt.Scanln(&aux)
					aux2, _ := strconv.Atoi(aux)
					if aux2 < 16 && aux2 > 0 {
						eleccionPiso3 += aux + ","
						i++
					} else {
						fmt.Println("Número no válido")
					}
				}
				sendMessage(fmt.Sprintf("%s:%s:%s", "JG", "3", eleccionPiso3))
				flag = false
			}
		case "2":
			//obtener monto
			sendQueryAmount()
		default:
			fmt.Println("Opción no válida, volviendo al menú de inicio")
		}
	}
}
