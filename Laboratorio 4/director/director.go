package main

import (
	"fmt"
	"os"
)

func main() {
	continuar := true
	for continuar {
		menuInicio := "\n     ====================\n    |1. Iniciar Mision   |\n    |2. Terminar programa|\n     ====================\nIngrese una opcion: "
		fmt.Println(menuInicio)

		var opcion string
		fmt.Scanln(&opcion)

		switch opcion {
		case "1":
			fmt.Println("Iniciando mision ...")
			menuOpciones := "\n     ============================\n    |1. Avanzar al siguiente piso|\n    |2. Volver al menu anterior  |\n     ============================\nIngrese una opcion: "
			fmt.Println(menuOpciones)

			var seguir string
			fmt.Scanln(&seguir)

			switch seguir {
			case "1":
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
