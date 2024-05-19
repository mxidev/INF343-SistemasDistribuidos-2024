package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	pb "grpc-server/munition"

	"google.golang.org/grpc"
)

func solicitarM(cliente pb.MunitionServiceClient, idequipos int, solAT int, solMP int) (bool, int) {

	//Realizar solicitud
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cliente.RequestMunition(ctx, &pb.MunitionRequest{IdGrupo: int32(idequipos), CantidadAt: int32(solAT), CantidadMp: int32(solMP)})

	//Problemas para ocupar la función
	if err != nil {
		fmt.Println("Error al ocupar función solicitarM()")
	}
	//Aprobado, retorna true y el numero del equipo a eliminar
	if res.Response {
		return true, idequipos
	} else {
		return false, idequipos
	}

}

func main() {
	// Espera 10 segundos.
	fmt.Println("Iniciando programa. Esperando 10 segundos para la primera consulta ...")
	duracion := 10 * time.Second
	time.Sleep(duracion)
	fmt.Println("Han pasado 10 segundos. Ejecutando primera consulta ...")

	// Inicializar la semilla para generar números realmente aleatorios.
	rand.Seed(time.Now().UnixNano())
	equipos := []int{1, 2, 3, 4}

	//Conectar al servidor tierra.go
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se ha podido conectar...")
	} else {
		cliente := pb.NewMunitionServiceClient(conn)

		//"Mientras haya equipos, sigue"
		for len(equipos) != 0 {

			equipoRand := rand.Intn(len(equipos))

			//"Mientras no estés satisfecho, sigue"
			for {

				municionAT := rand.Intn(11) + 20
				municionMP := rand.Intn(6) + 10

				termino, eliminar := solicitarM(cliente, equipoRand, municionAT, municionMP)
				if termino {
					fmt.Printf("Solicitando %d AT y %d MP ; Resolucion: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicacion con equipo %d\n", municionAT, municionMP, equipos[equipoRand])
					equipos = append(equipos[:eliminar], equipos[eliminar+1:]...) // Se elimina el equipo que termina la comunicacion para no volver a iterarlo
					break
				} else {
					fmt.Printf("Solicitando %d AT y %d MP ; Resolucion: -- DENEGADA -- ; Reintentando en 3 segs... \n", municionAT, municionMP)
					time.Sleep(3 * time.Second)
				}
			}
		}

	}

}
