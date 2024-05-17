package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func comunicar(decision int) {
	//Realizar solicitud
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//poner el pb

	//Problemas para ocupar la función
	if err != nil {
		fmt.Println("Error al ocupar función solicitarM()")
	}
}

func main() {
	var i int
	begin := true

	//
	for begin {
		fmt.Print("Si está listo ingrese [1]\n")
		fmt.Scanf("%d", &i)
		if i == 1 {
			//mandar que está listo
			begin = false
		}
	}
	flag := true
	estado := true
	contPiso := 0
	for flag {
		var eleccion string

		//Caso en el que el jugador haya muerto
		if estado == false {
			fmt.Print("Usted ha muerto...")
			goto muerto
		}

		//Opciones para el jugador
		MenuInicio := "Jugador, ¿Qué desea hacer?\n [1] Explorar piso\n [2] Ver monto acumulado\nIngrese opción:"
		fmt.Print(MenuInicio)

		fmt.Scanln(&eleccion)

		switch eleccion {
		case "1":
			var eleccionPiso string
			var eleccionPiso3 [5]string

			//Piso 1
			if contPiso == 0 {
				flagPiso1 := true
				for flagPiso1 {
					fmt.Print("Jugador, ¿Qué desea ocupar?\n [1] Escopeta\n [2] Rifle automático\n [3] Puños eléctricos ")
					fmt.Scanln(&eleccionPiso)

					switch eleccionPiso {
					case "1":
						//comunicar 1
						contPiso++
						flagPiso1 = false

					case "2":
						//comunicar 2
						contPiso++
						flagPiso1 = false
					case "3":
						//comunicar 3
						contPiso++
						flagPiso1 = false
					default:
						fmt.Println("Opción no válida")
					}
				}

				//Piso 2
			} else if contPiso == 1 {
				flagPiso2 := true
				for flagPiso2 {
					fmt.Print("Jugador, ¿Dónde desea ir?\n [1] A\n [2] B ")
					fmt.Scanln(&eleccionPiso)

					switch eleccionPiso {
					case "1":
						//comunicar A
						contPiso++
						flagPiso2 = false

					case "2":
						//comunicar B
						contPiso++
						flagPiso2 = false
					default:
						fmt.Println("Opción no válida")
					}
				}

				//Piso 2
			} else if contPiso == 2 {
				i := 0
				fmt.Print("A continuación debe elegir 5 números del 1 al 15\n")
				for i < 5 {
					var aux string
					fmt.Printf("Elija el número %d: ", i+1)
					fmt.Scanln(&aux)
					aux2, _ := strconv.Atoi(aux)
					if aux2 < 16 && aux2 > 0 {
						eleccionPiso3[i] = aux
						i++
					} else {
						fmt.Println("Número no válido")
					}

				}
				flag = false
			}

		case "2":
			//obtener monto

		default:
			fmt.Println("Opción no válida, volviendo al menú de inicio")
		}

	}

muerto:
}
