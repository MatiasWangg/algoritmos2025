package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	sistema "tp2/sistemaVuelos"
	"time"
)
const _LAYOUT = "2006-01-02T15:04:05"

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
			return fmt.Errorf("Error en comando agregar_archivo")
		}
		archivo := partes[1]

		return sistemaVuelos.AgregarArchivo(archivo)


	case "ver_tablero":
		if len(partes) != 5 {
			return fmt.Errorf("Error en comando ver_tablero")
		}

		cantidadDeVuelos, err := strconv.Atoi(partes[1])
		if err != nil {
			return fmt.Errorf("Error en comando prioridad_vuelos")
		}
		modo := partes[2]
		fechaDesde := partes[3]
		fechaHasta := partes[4]
		return sistemaVuelos.VerTablero(cantidadDeVuelos, modo, fechaDesde, fechaHasta)


	case "info_vuelo":
		if len(partes) != 2 {
			return fmt.Errorf("Error en comando info_vuelo")
		}
		codigo := partes[1]
		return sistemaVuelos.InfoVuelo(codigo)


	case "prioridad_vuelos":
		if len(partes) != 2 {
			return fmt.Errorf("Error en comando prioridad_vuelos")
		}
		n, err := strconv.Atoi(partes[1])
		if n <= 0 {
			return fmt.Errorf("Error en comando prioridad_vuelos")
		}

		if err != nil {
			return fmt.Errorf("Error en comando prioridad_vuelos")
		}
		return sistemaVuelos.Prioridad_vuelos(n)


	case "siguiente_vuelo":
		if len(partes) != 4 {
		return fmt.Errorf("Error en comando siguiente_vuelo")
		}

		origen := partes[1]
		destino := partes[2]
		fecha, err := time.Parse(_LAYOUT, partes[3])
		if err != nil {
			return fmt.Errorf("Error en comando siguiente_vuelo")
		}

		return sistemaVuelos.SiguienteVuelo(origen, destino, fecha)

	case "borrar":
		if len(partes)!= 3{
			return fmt.Errorf("Error en comando borrar")
		}
		fechaDesde := partes[1]
		fechaHasta := partes[2]
		return sistemaVuelos.Borrar(fechaDesde,fechaHasta)
	default:
		return fmt.Errorf("Error en comando %s", partes[0])
	}
}

func imprimirResultado(resultado error) {
	if resultado != nil {
		fmt.Fprintf(os.Stderr, "%s\n", resultado)
	} else {
		fmt.Println("OK")
	}
}
