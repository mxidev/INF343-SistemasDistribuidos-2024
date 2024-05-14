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

	// Servidor central
	addr := "127.0.0.1:3000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error con la conexion al servidor central:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// arreglo para mensaje a enviar al servidor central
	var mensaje [4]string

	pan_c1 := rand.Intn(10) + 7
	leche_c1 := rand.Intn(10) + 7
	agua_c1 := rand.Intn(10) + 7

	pan_c2 := rand.Intn(10) + 7
	leche_c2 := rand.Intn(10) + 7
	agua_c2 := rand.Intn(10) + 7

	pan_c3 := rand.Intn(10) + 7
	leche_c3 := rand.Intn(10) + 7
	agua_c3 := rand.Intn(10) + 7

	array_capitanes := [3]string{"C1", "C2", "C3"}

	continuar := true
	contador := 10
	for continuar {
		opt := rand.Intn(3)
		capitan := array_capitanes[opt]
		if capitan == "C1" {
			mensaje[0] = "1" // # Cap.
			fmt.Println("Nave de C1")
			pan_c1 -= rand.Intn(2) + 2
			leche_c1 -= rand.Intn(2) + 2
			agua_c1 -= rand.Intn(2) + 2
			if pan_c1 < 7 {
				fmt.Printf("%s tiene %d cantidad de pan. Enviando solicitud ...\n", capitan, pan_c1)
				mensaje[1] = "1"                  // Tipo sol.
				mensaje[2] = "7"                  // Cantidad Sol.
				mensaje[3] = strconv.Itoa(pan_c1) // Restante C3
			} else if leche_c1 < 7 {
				fmt.Printf("%s tiene %d cantidad de leche. Enviando solicitud ...\n", capitan, leche_c1)
				mensaje[1] = "2"                    // Tipo sol.
				mensaje[2] = "7"                    // Cantidad Sol.
				mensaje[3] = strconv.Itoa(leche_c1) // Restante C3
			} else if agua_c1 < 7 {
				fmt.Printf("%s tiene %d cantidad de agua. Enviando solicitud ...\n", capitan, agua_c1)
				mensaje[1] = "3"                   // Tipo sol.
				mensaje[2] = "7"                   // Cantidad Sol.
				mensaje[3] = strconv.Itoa(agua_c1) // Restante C3
			}
		} else if capitan == "C2" {
			mensaje[0] = "2"
			fmt.Println("Nave de C2")
			pan_c2 -= rand.Intn(2) + 2
			leche_c2 -= rand.Intn(2) + 2
			agua_c2 -= rand.Intn(2) + 2
			if pan_c2 < 7 {
				fmt.Printf("%s tiene %d cantidad de pan. Enviando solicitud ...\n", capitan, pan_c2)
				mensaje[1] = "1"                  // Tipo sol.
				mensaje[2] = "7"                  // Cantidad Sol.
				mensaje[3] = strconv.Itoa(pan_c2) // Restante C3
			} else if leche_c2 < 7 {
				fmt.Printf("%s tiene %d cantidad de leche. Enviando solicitud ...\n", capitan, leche_c2)
				mensaje[1] = "2"                    // Tipo sol.
				mensaje[2] = "7"                    // Cantidad Sol.
				mensaje[3] = strconv.Itoa(leche_c2) // Restante C3
			} else if agua_c2 < 7 {
				fmt.Printf("%s tiene %d cantidad de agua. Enviando solicitud ...\n", capitan, agua_c2)
				mensaje[1] = "3"                   // Tipo sol.
				mensaje[2] = "7"                   // Cantidad Sol.
				mensaje[3] = strconv.Itoa(agua_c2) // Restante C3
			}
		} else {
			fmt.Println("Nave de C3")
			mensaje[0] = "3"
			pan_c3 -= rand.Intn(2) + 2
			leche_c3 -= rand.Intn(2) + 2
			agua_c3 -= rand.Intn(2) + 2
			if pan_c3 < 7 {
				fmt.Printf("%s tiene %d cantidad de pan. Enviando solicitud ...\n", capitan, pan_c3)
				mensaje[1] = "1"                  // Tipo sol.
				mensaje[2] = "7"                  // Cantidad Sol.
				mensaje[3] = strconv.Itoa(pan_c3) // Restante C3
			} else if leche_c3 < 7 {
				fmt.Printf("%s tiene %d cantidad de leche. Enviando solicitud ...\n", capitan, leche_c3)
				mensaje[1] = "2"                    // Tipo sol.
				mensaje[2] = "7"                    // Cantidad Sol.
				mensaje[3] = strconv.Itoa(leche_c3) // Restante C3
			} else if agua_c3 < 7 {
				fmt.Printf("%s tiene %d cantidad de agua. Enviando solicitud ...\n", capitan, agua_c3)
				mensaje[1] = "3"                   // Tipo sol.
				mensaje[2] = "7"                   // Cantidad Sol.
				mensaje[3] = strconv.Itoa(agua_c3) // Restante C3
			}
		}

		// Codigo para enviar el mensaje al servidor central
		enviar := mensaje[0] + " " + mensaje[1] + " " + mensaje[2] + " " + mensaje[3] + "\n"
		msg := []byte(enviar)
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("Error -> ", err)
			break
		}

		// Codigo para la respuesta que entregara servidor central
		resp := make([]byte, 1024)
		index, err := conn.Read(resp)
		if err != nil {
			fmt.Println("Error con recibir respuesta del servidor central: ", err)
			break
		}

		// Procesamiento del mensaje recibido por Servidor Central
		fmt.Println("Respuesta recibida por parte del Servidor Central")
		format := strings.TrimSpace(string(resp[:index]))
		slices := strings.Split(format, " ")
		numero_capitan, _ := strconv.Atoi(slices[0])
		cantidad_solicitud, _ := strconv.Atoi(slices[1])

		if numero_capitan == 1 {
			if mensaje[1] == "1" {
				pan_c1 += cantidad_solicitud
			} else if mensaje[1] == "2" {
				leche_c1 += cantidad_solicitud
			} else if mensaje[1] == "3" {
				agua_c1 += cantidad_solicitud
			}

		} else if numero_capitan == 2 {
			if mensaje[1] == "1" {
				pan_c2 += cantidad_solicitud
			} else if mensaje[1] == "2" {
				leche_c2 += cantidad_solicitud
			} else if mensaje[1] == "3" {
				agua_c2 += cantidad_solicitud
			}

		} else if numero_capitan == 3 {
			if mensaje[1] == "1" {
				pan_c3 += cantidad_solicitud
			} else if mensaje[1] == "2" {
				leche_c3 += cantidad_solicitud
			} else if mensaje[1] == "3" {
				agua_c3 += cantidad_solicitud
			}
		}

		contador -= 1
		fmt.Println("Check!")
		if contador == 0 {
			continuar = false
		}
		enviar = ""
	}
}
