package sistemavuelos

import (
	"tdas/diccionario"
	"tp2/vuelo"
)

type Sistema struct {
	vuelos diccionario.Diccionario[string, *vuelo.Vuelo] //string (clave) es codigo del vuelo
	//La clave y valor del abb puede ser cambiados, puse estos porque pense que era la mejor idea
	vuelosABB diccionario.DiccionarioOrdenado[string, *vuelo.Vuelo] //string (clave) es el horario del vuelo
}

func CrearSistema() *Sistema {
	sistema := new(Sistema)
	sistema.vuelos = diccionario.CrearHash[string, *vuelo.Vuelo]()
	sistema.vuelosABB = diccionario.CrearABB[string, *vuelo.Vuelo](func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	return sistema
}
