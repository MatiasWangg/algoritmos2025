package sistemavuelos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"time"
	"tp2/vuelo"
)

const _LAYOUT = "2006-01-02T15:04:05"

func (s *Sistema) AgregarArchivo(archivo string) error {
	contenido, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("Error en comando agregar_archivo")
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

		fecha, err := time.Parse(_LAYOUT, datos[6])
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
		if s.vuelos.Pertenece(datos[0]) {
			fechaAntigua := s.vuelos.Obtener(datos[0]).Fecha
			fechaStr := fechaAntigua.Format(_LAYOUT)
			claveAntigua := fmt.Sprintf("%s|%s", fechaStr, datos[0])
			s.vuelosABB.Borrar(claveAntigua)
		}

		s.vuelos.Guardar(datos[0], vuelo)

		fechaStr := fecha.Format(_LAYOUT)
		claveABB := fmt.Sprintf("%s|%s", fechaStr, datos[0])
		s.vuelosABB.Guardar(claveABB, vuelo)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error al leer archivo")
	}

	return nil
}

func (s *Sistema) VerTablero(cantidadDeVuelos int, modo, fechaDesde, fechaHasta string) error {
	if cantidadDeVuelos <= 0 {
		return fmt.Errorf("Error en comando ver_tablero")
	} else if modo != "asc" && modo != "desc" {
		return fmt.Errorf("Error en comando ver_tablero")
	} else if fechaHasta < fechaDesde {
		return fmt.Errorf("Error en comando ver_tablero")
	}

	claveInicio := fechaDesde + "|"
	claveFin := fechaHasta + "|~"
	iteradorRango := s.vuelosABB.IteradorRango(&claveInicio, &claveFin)
	arrayVuelosOrdenados := []string{}

	for iteradorRango.HaySiguiente() {
		_, vueloActual := iteradorRango.VerActual()
		fechaStr := vueloActual.ObtenerFecha().Format(_LAYOUT)
		infoRes := fmt.Sprintf("%s - %s\n", fechaStr, vueloActual.ObtenerCodigo())
		arrayVuelosOrdenados = append(arrayVuelosOrdenados, infoRes)
		iteradorRango.Siguiente()
	}
	if modo == "asc" {
		for i := 0; i < cantidadDeVuelos && i < len(arrayVuelosOrdenados); i++ {
			fmt.Printf(arrayVuelosOrdenados[i])
		}
	} else {
		for i := len(arrayVuelosOrdenados) - 1; i >= len(arrayVuelosOrdenados)-cantidadDeVuelos && i >= 0; i-- {
			fmt.Printf(arrayVuelosOrdenados[i])
		}
	}
	return nil
}

func (s *Sistema) Borrar(fechaDesde, fechaHasta string) error {
	if fechaHasta < fechaDesde {
		return fmt.Errorf("Error en comando borrar")
	}
	claveInicio := fechaDesde + "|"
	claveFin := fechaHasta + "|~"
	iteradorRango := s.vuelosABB.IteradorRango(&claveInicio, &claveFin)
	clavesABBAeliminar := []string{}

	for iteradorRango.HaySiguiente() {
		claveActual, vueloActual := iteradorRango.VerActual()
		vueloEliminado := s.vuelos.Borrar(vueloActual.ObtenerCodigo())
		clavesABBAeliminar = append(clavesABBAeliminar, claveActual)
		imprimirVuelo(vueloEliminado)
		iteradorRango.Siguiente()
	}
	for _, clave := range clavesABBAeliminar {
		s.vuelosABB.Borrar(clave)
	}
	return nil
}

func (s *Sistema) InfoVuelo(codigo string) error {
	if !s.vuelos.Pertenece(codigo) {
		return fmt.Errorf("Error en comando info_vuelo")
	}

	vuelo := s.vuelos.Obtener(codigo)

	imprimirVuelo(vuelo)

	return nil
}

func (s *Sistema) Prioridad_vuelos(k int) error {
	vuelos := make([]*vuelo.Vuelo, 0, s.vuelos.Cantidad())

	s.vuelos.Iterar(func(codigo string, vuelo *vuelo.Vuelo) bool {
		vuelos = append(vuelos, vuelo)
		return true
	})

	cmp := func(a, b *vuelo.Vuelo) int {
		if a.ObtenerPrioridad() != b.ObtenerPrioridad() {
			return a.ObtenerPrioridad() - b.ObtenerPrioridad()
		}
		if a.ObtenerCodigo() < b.ObtenerCodigo() {
			return 1
		} else if a.ObtenerCodigo() > b.ObtenerCodigo() {
			return -1
		}
		return 0
	}

	heap := cola_prioridad.CrearHeapArr(vuelos, cmp)

	for i := 0; i < k && !heap.EstaVacia(); i++ {
		vuelo := heap.Desencolar()
		fmt.Printf("%d - %s\n", vuelo.ObtenerPrioridad(), vuelo.ObtenerCodigo())
	}

	return nil
}

func (s *Sistema) SiguienteVuelo(origen, destino string, fecha time.Time) error {
	fechaStr := fecha.Format(_LAYOUT)
	iterador := s.vuelosABB.IteradorRango(&fechaStr, nil)

	for iterador.HaySiguiente() {
		_, vuelo := iterador.VerActual()
		if vuelo.ObtenerOrigen() == origen && vuelo.ObtenerDestino() == destino {
			s.InfoVuelo(vuelo.ObtenerCodigo())
			return nil
		}
		iterador.Siguiente()
	}

	fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fechaStr)
	return nil
}

func imprimirVuelo(v *vuelo.Vuelo) {
	canceladoInt := v.EstaCanceladoInt()
	fechaStr := v.ObtenerFecha().Format(_LAYOUT)

	fmt.Printf("%s %s %s %s %s %d %s %d %d %d\n",
		v.ObtenerCodigo(),
		v.ObtenerAerolinea(),
		v.ObtenerOrigen(),
		v.ObtenerDestino(),
		v.ObtenerMatricula(),
		v.ObtenerPrioridad(),
		fechaStr,
		v.ObtenerRetrasoSalida(),
		v.ObtenerTiempoVuelo(),
		canceladoInt,
	)
}
