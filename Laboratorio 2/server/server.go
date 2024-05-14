package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Se configura la conexion con el puerto
	puerto := ":8080"
	listn, err := net.Listen("tcp", puerto)
	if err != nil {
		fmt.Println("Error creando socket: ", err)
		os.Exit(1)
	}
	defer listn.Close()
	fmt.Println("Servidor Central operativo!")

	// Gestion de mensajes entrantes
	for {
		conn, err := listn.Accept()
		if err != nil {
			fmt.Println("Error con la conexion: ", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	terminar := true
	pan := rand.Intn(30) + 14
	leche := rand.Intn(30) + 14
	agua := rand.Intn(30) + 14
	fmt.Println("Suministros del Servidor Central:")
	fmt.Println("	> pan actual: ", pan)
	fmt.Println("	> leche actual: ", leche)
	fmt.Println("	> agua actual: ", agua)

	var respuesta [2]string

	for terminar {
		// Mensaje entrante
		mensaje := make([]byte, 1024)
		index, err := conn.Read(mensaje)
		if err != nil {
			fmt.Println("Error con lectura: ", err)
			return
		}

		// Procesamiento de mensaje entrante
		format := strings.TrimSpace(string(mensaje[:index]))
		slices := strings.Split(format, " ")
		numero_capitan := slices[0]
		tipo_solicitud := slices[1]
		cantidad_solicitud, _ := strconv.Atoi(slices[2])
		restante := slices[3]

		if tipo_solicitud == "1" { //1=pan
			fmt.Printf("\nRecepción de solicitud del capitán C%s, solicita %d unidades de pan y le quedan %s unidades.", numero_capitan, cantidad_solicitud, restante)
			if cantidad_solicitud <= pan {
				pan = pan - cantidad_solicitud
				respuesta[0] = numero_capitan
				respuesta[1] = strconv.Itoa(cantidad_solicitud)
				fmt.Printf("\nSuministros pan asignado al capitán C%s, ahora quedan %d panes disponibles para entregar", numero_capitan, pan)

			} else {
				respuesta[0] = numero_capitan
				respuesta[1] = "0"
				fmt.Println("\n\nNo hay suficientes suministros, terminando programa..")
				terminar = false
			}
		} else if tipo_solicitud == "2" { //2=leche
			fmt.Printf("\nRecepción de solicitud del capitán C%s, solicita %d unidades de leche y le quedan %s unidades.", numero_capitan, cantidad_solicitud, restante)
			if cantidad_solicitud <= leche {
				respuesta[0] = numero_capitan
				respuesta[1] = strconv.Itoa(cantidad_solicitud)
				leche = leche - cantidad_solicitud
				fmt.Printf("\nSuministros de leche asignado al capitán C%s, ahora quedan %d leches disponibles para entregar", numero_capitan, leche)

			} else {
				respuesta[0] = numero_capitan
				respuesta[1] = "0"
				fmt.Println("\n\nNo hay suficientes suministros, terminando programa..")
				terminar = false
			}

		} else if tipo_solicitud == "3" { //3=agua
			fmt.Printf("\nRecepción de solicitud del capitán C%s, solicita %d unidades de agua y le quedan %s unidades.", numero_capitan, cantidad_solicitud, restante)
			if cantidad_solicitud <= agua {
				respuesta[0] = numero_capitan
				respuesta[1] = strconv.Itoa(cantidad_solicitud)
				agua = agua - cantidad_solicitud
				fmt.Printf("\nSuministros de agua asignado al capitán C%s, ahora quedan %d unidades de agua disponibles para entregar", numero_capitan, agua)
			} else {
				respuesta[0] = numero_capitan
				respuesta[1] = "0"
				fmt.Println("\n\nNo hay suficientes suministros, terminando programa..")
				terminar = false
			}
		}

		fmt.Println("\nEnviando respuesta al capitán")
		env := respuesta[0] + " " + respuesta[1] + "\n"
		msg := []byte(env)
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("Error -> ", err)
			break
		}

		fmt.Println("\n	> pan actual: ", pan)
		fmt.Println("	> leche actual: ", leche)
		fmt.Println("	> agua actual: ", agua)
	}

	// Se cierra la conexion al terminar la iteracion
	defer conn.Close()
}
