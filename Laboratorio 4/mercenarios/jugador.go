package main

import "fmt"

func main() {
	var i int
	begin := true

	for begin {
		fmt.Print("Si está listo ingrese 1\n")
		fmt.Scanf("%d", &i)
		if i == 1 {
			//mandar que está listo
			begin = false
		}
	}
}
