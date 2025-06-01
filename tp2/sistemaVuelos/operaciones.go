package sistemavuelos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"tp2/vuelo"
)

const LAYOUT = "2006-01-02T15:04:05"

//Lugar donde se van a implementar las firmas de primitivas

func (s *Sistema) AgregarArchivo(archivo string) error {
	contenido, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo")
	}
	defer contenido.Close()

	scanner := bufio.NewScanner(contenido)

	for scanner.Scan() {
		datos := strings.Split(scanner.Text(), ",")
		if len(datos) != 10 {
			return fmt.Errorf("formato inválido en la línea del archivo")
		}

		prioridad, err := strconv.Atoi(datos[5])
		if err != nil {
			return fmt.Errorf("error al parsear prioridad")
		}

		fecha, err := time.Parse(LAYOUT, datos[6])
		if err != nil {
			return fmt.Errorf("error al parsear fecha")
		}

		retraso, err := strconv.Atoi(datos[7])
		if err != nil {
			return fmt.Errorf("error al parsear retraso")
		}

		tiempo, err := strconv.Atoi(datos[8])
		if err != nil {
			return fmt.Errorf("error al parsear tiempo")
		}

		canceladoInt, err := strconv.Atoi(datos[9])
		if err != nil {
			return fmt.Errorf("error al parsear cancelado")
		}
		cancelado := canceladoInt != 0

		vuelo := vuelo.CrearVuelo(
			datos[0], datos[1], datos[2], datos[3], datos[4],
			prioridad, retraso, tiempo, fecha, cancelado,
		)

		//Se ingresa el vuelo al Hash
		s.vuelos.Guardar(datos[0], vuelo)

		//Se ingresa el vuelo al ABB
		fechaStr := fecha.Format(LAYOUT)
		s.vuelosABB.Guardar(fechaStr, vuelo)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error al leer archivo")
	}

	return nil
}


func (s *Sistema) InfoVuelo(codigo string) error {
	if !s.vuelos.Pertenece(codigo) {
		return fmt.Errorf("vuelo no encontrado") 
	}

	vuelo:= s.vuelos.Obtener(codigo)

	canceladoInt := vuelo.EstaCanceladoInt()

	fechaStr := vuelo.Fecha.Format(LAYOUT)

	fmt.Printf("%s %s %s %s %s %d %s %02d %d %d\n",
		vuelo.Codigo,
		vuelo.Aerolinea,
		vuelo.Origen,
		vuelo.Destino,
		vuelo.Matricula,
		vuelo.Prioridad,
		fechaStr,
		vuelo.RetrasoSalida,
		vuelo.TiempoVuelo,
		canceladoInt,
	)

	return nil
}

