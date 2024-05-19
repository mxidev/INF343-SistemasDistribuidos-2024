package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Valores al azar para los bots
	rand.Seed(time.Now().UnixNano())
	piso1 := rand.Intn(3) + 1
	piso2 := rand.Intn(2) + 1
	piso3_1 := rand.Intn(15) + 1
	piso3_2 := rand.Intn(15) + 1
	piso3_3 := rand.Intn(15) + 1
	piso3_4 := rand.Intn(15) + 1
	piso3_5 := rand.Intn(15) + 1

	fmt.Println(piso1, piso2, piso3_1, piso3_2)

	//El mercenario murió
	if piso3_2 < 8 {
		fmt.Println("Mercenario 1 has muerto..")
		goto muerto
	}
	fmt.Println(piso3_3, piso3_4, piso3_5)
	//mandar que está listo
	//mandar decisiones

muerto:
}
