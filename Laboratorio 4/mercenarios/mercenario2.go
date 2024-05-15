package main

import (
	"fmt"
	"math/rand"
)

func main() {

	piso1 := rand.Intn(3) + 1
	piso2 := rand.Intn(2) + 1
	piso3_1 := rand.Intn(15) + 1
	piso3_2 := rand.Intn(15) + 1
	piso3_3 := rand.Intn(15) + 1
	piso3_4 := rand.Intn(15) + 1
	piso3_5 := rand.Intn(15) + 1

	fmt.Println(piso1, piso2, piso3_1, piso3_2)

	if piso3_2 < 8 {
		fmt.Println("Mercenario 2 has muerto..")
		goto muerto
	}
	fmt.Println(piso3_3, piso3_4, piso3_5)
	//mandar que está listo

muerto:
}
