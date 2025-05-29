package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
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

func procesarComando(comando string, sistemaVuelos *sistema.Sistema) error{
	partes := strings.Fields(comando)
	
	if len(partes) == 0 {
		return fmt.Errorf("Comando no reconocido")
	}
	
	switch partes[0] {
	case "agregar_archivo":
		
	case "ver_tablero":
	
	case "info_vuelo":

	case "prioridad_vuelos":
	
	case "siguiente_vuelo":
	
	case "borrar":

	default:
		return fmt.Errorf("Comando no reconocido")
	}
}

func imprimirResultado(resultado error) {
	if resultado != nil {
		fmt.Fprintf(os.Stderr, "Error en el comando %s \n", resultado)
	}else {
		fmt.Println("OK")
	}
}