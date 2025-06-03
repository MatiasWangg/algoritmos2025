package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	sistema "tp2/sistemaVuelos"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sistemaVuelos := sistema.CrearSistema()

	for scanner.Scan() {
		comando := scanner.Text()

		resultado := procesarComando(comando, sistemaVuelos)
		imprimirResultado(resultado)
	}
}

func procesarComando(comando string, sistemaVuelos *sistema.Sistema) error {
	partes := strings.Fields(comando)

	if len(partes) == 0 {
		return fmt.Errorf("comando no reconocido")
	}

	switch partes[0] {
	case "agregar_archivo":
		if len(partes) != 2 {
			return fmt.Errorf("error en agregar archivo")
		}
		archivo := partes[1]

		return sistemaVuelos.AgregarArchivo(archivo)

	case "ver_tablero":
		if len(partes) != 5 {
			return fmt.Errorf("error en ver tablero")
		}

		cantidadDeVuelos, _ := strconv.Atoi(partes[1])
		modo := partes[2]
		fechaDesde := partes[3]
		fechaHasta := partes[4]
		return sistemaVuelos.VerTablero(cantidadDeVuelos, modo, fechaDesde, fechaHasta)
	case "info_vuelo":
		if len(partes) != 2 {
			return fmt.Errorf("error en info_vuelo")
		}
		codigo := partes[1]
		return sistemaVuelos.InfoVuelo(codigo)
	case "prioridad_vuelos":
		if len(partes) != 2 {
		return fmt.Errorf("error en prioridad_vuelos")
		}
		n, err := strconv.Atoi(partes[1])
		if err != nil {
			return fmt.Errorf("error en prioridad_vuelos")
		}
		return sistemaVuelos.Prioridad_vuelos(n)
	case "siguiente_vuelo":
		return nil
	case "borrar":
		return nil
	default:
		return fmt.Errorf("comando no reconocido")
	}
}

func imprimirResultado(resultado error) {
	if resultado != nil {
		fmt.Fprintf(os.Stderr, "Error en el comando %s \n", resultado)
	} else {
		fmt.Println("OK")
	}
}
