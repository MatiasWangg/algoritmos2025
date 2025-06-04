package sistemavuelos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"tp2/vuelo"
	"tdas/cola_prioridad"
)

const LAYOUT = "2006-01-02T15:04:05"

//Lugar donde se van a implementar las firmas de primitivas

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
		claveABB := fmt.Sprintf("%s|%s", fechaStr, datos[0]) // fecha + código
		s.vuelosABB.Guardar(claveABB, vuelo)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error al leer archivo")
	}

	return nil
}

func (s *Sistema) VerTablero(cantidadDeVuelos int, modo , fechaDesde, fechaHasta string)error {
	if cantidadDeVuelos <= 0  {
		return fmt.Errorf("La cantidad de vuelos pedidos no puede ser menor o igual a 0") 
	}else if (modo != "asc" && modo != "desc"){
		return fmt.Errorf("Error en especificar el modo en que desea la tabla de vuelos") 
	}else if fechaHasta < fechaDesde{
		return fmt.Errorf("Error en especificar el rango de fechas en que se pide la tabla") 
	}	

	claveInicio := fechaDesde + "|"
	claveFin := fechaHasta + "|~" // ~ es para que tome todos los vuelos de ese día
	iteradorRango := s.vuelosABB.IteradorRango(&claveInicio,&claveFin)
	arrayVuelosOrdenados := []string{}
	if !iteradorRango.HaySiguiente(){
		fmt.Printf("No hay vuelos dentro de ese rango de fechas\n")
	}
	for iteradorRango.HaySiguiente(){
		_, vueloActual := iteradorRango.VerActual()
		fechaStr := vueloActual.Fecha.Format(LAYOUT)
		infoRes := fmt.Sprintf("%s - %s\n", fechaStr, vueloActual.Codigo)
		arrayVuelosOrdenados = append(arrayVuelosOrdenados, infoRes)
		iteradorRango.Siguiente()
	}
	if modo == "asc"{
		for i := 0; i< cantidadDeVuelos && i <len(arrayVuelosOrdenados); i++{
			fmt.Printf(arrayVuelosOrdenados[i])
		}
	}else{
		for i := len(arrayVuelosOrdenados)-1; i>= len(arrayVuelosOrdenados) - cantidadDeVuelos; i--{
			fmt.Printf(arrayVuelosOrdenados[i])
		}
	}
	return nil
}

func (s *Sistema) Borrar(fechaDesde, fechaHasta string)error {
	if fechaHasta < fechaDesde{
		return fmt.Errorf("Error en especificar el rango de fechas en que se quiere eliminar") 
	}	
	claveInicio := fechaDesde + "|"
	claveFin := fechaHasta + "|~"
	iteradorRango := s.vuelosABB.IteradorRango(&claveInicio,&claveFin)
	if !iteradorRango.HaySiguiente(){
		fmt.Printf("No hay vuelos dentro de ese rango de fechas\n")
	}
	clavesABBAeliminar := []string{}

	for iteradorRango.HaySiguiente(){
		claveActual, vueloActual := iteradorRango.VerActual()
		vueloEliminado := s.vuelos.Borrar(vueloActual.Codigo)
		clavesABBAeliminar = append(clavesABBAeliminar, claveActual)
		fechaStr := vueloEliminado.Fecha.Format(LAYOUT)
		canceladoInt := vueloEliminado.EstaCanceladoInt()
		fmt.Printf("%s %s %s %s %s %d %s %d %d %d\n",
			vueloEliminado.Codigo,
			vueloEliminado.Aerolinea,
			vueloEliminado.Origen,
			vueloEliminado.Destino,
			vueloEliminado.Matricula,
			vueloEliminado.Prioridad,
			fechaStr,
			vueloEliminado.RetrasoSalida,
			vueloEliminado.TiempoVuelo,
			canceladoInt,
		)
		iteradorRango.Siguiente()
	}
	for _, clave := range(clavesABBAeliminar){
		s.vuelosABB.Borrar(clave)
	}
	return nil
}


func (s *Sistema) InfoVuelo(codigo string) error {
	if !s.vuelos.Pertenece(codigo) {
		return fmt.Errorf("Error en comando info_vuelo") 
	}

	vuelo:= s.vuelos.Obtener(codigo)

	canceladoInt := vuelo.EstaCanceladoInt()

	fechaStr := vuelo.Fecha.Format(LAYOUT)

	fmt.Printf("%s %s %s %s %s %d %s %d %d %d\n",
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


func (s *Sistema) Prioridad_vuelos(k int) error {
	vuelos := make([]*vuelo.Vuelo, 0, s.vuelos.Cantidad())

	s.vuelos.Iterar(func(codigo string, vuelo *vuelo.Vuelo) bool {
		vuelos = append(vuelos, vuelo)
		return true
	})

	heap := cola_prioridad.CrearHeapArr(vuelos, func(a, b *vuelo.Vuelo) int {
		if a.Prioridad != b.Prioridad {
			return a.Prioridad - b.Prioridad 
		}
		if a.Codigo < b.Codigo {
			return 1
		} else if a.Codigo > b.Codigo {
			return -1
		}
		return 0
	})

	for i := 0; i < k && !heap.EstaVacia(); i++ {
		vuelo := heap.Desencolar()
		fmt.Printf("%d - %s\n", vuelo.Prioridad, vuelo.Codigo)
	}
	
	return nil
}

