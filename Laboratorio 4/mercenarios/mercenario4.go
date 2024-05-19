package main

import (
	"fmt"
	pb "grpc/proto"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

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

func main() {

	//Valores al azar para los bots
	rand.Seed(time.Now().UnixNano())
	piso1 := rand.Intn(3) + 1
	piso2 := rand.Intn(2) + 1
	piso3_1 := rand.Intn(15) + 1
	piso3_2 := rand.Intn(15) + 1
	for piso3_2 == piso3_1 {
		piso3_2 = rand.Intn(15) + 1
	}
	piso3_3 := rand.Intn(15) + 1
	for piso3_3 == piso3_1 || piso3_3 == piso3_2 {
		piso3_3 = rand.Intn(15) + 1
	}
	piso3_4 := rand.Intn(15) + 1
	for piso3_4 == piso3_1 || piso3_4 == piso3_2 || piso3_4 == piso3_3 {
		piso3_4 = rand.Intn(15) + 1
	}
	piso3_5 := rand.Intn(15) + 1
	for piso3_5 == piso3_1 || piso3_5 == piso3_2 || piso3_5 == piso3_3 || piso3_5 == piso3_4 {
		piso3_5 = rand.Intn(15) + 1
	}

	probs := [5]string{strconv.Itoa(piso3_1), strconv.Itoa(piso3_2), strconv.Itoa(piso3_3), strconv.Itoa(piso3_4), strconv.Itoa(piso3_5)}
	elecciones := ""

	i := 0
	for i < 5 {
		elecciones += probs[i] + ","
		i++
	}

	fmt.Println(elecciones)

	//fmt.Println(piso1, piso2, piso3_1, piso3_2)

	// 	//El mercenario murió
	// 	if piso3_2 < 8 {
	// 		fmt.Println("Mercenario 1 has muerto..")
	// 		goto muerto
	// 	}
	// 	fmt.Println(piso3_3, piso3_4, piso3_5)
	// 	//mandar que está listo
	// 	//mandar decisiones

	// muerto:
	sendMessage(fmt.Sprintf("%s:%s:%s", "M4", "1", strconv.Itoa(piso1)))
	sendMessage(fmt.Sprintf("%s:%s:%s", "M4", "2", strconv.Itoa(piso2)))
	sendMessage(fmt.Sprintf("%s:%s:%s", "M4", "3", elecciones))
}
