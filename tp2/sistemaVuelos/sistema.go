package sistemavuelos

import (
	"tp2/vuelo"
	"tdas/diccionario"
)

type Sistema struct {
	vuelos diccionario.Diccionario[string, *vuelo.Vuelo] //string (clave) es codigo del vuelo
}

func CrearSistema() *Sistema {
	sistema := new(Sistema)
	sistema.vuelos = diccionario.CrearHash[string, *vuelo.Vuelo]()
	return sistema
}

//Firmas de primitivas provisorias 
// func (s *Sistema) AgregarArchivo() 

// func (s *Sistema) VerTablero() 

// func (s *Sistema) InfoVuelo() 

// func (s *Sistema) PrioridadVuelos() 

// func (s *Sistema) SiguienteVuelo() ()

// func (s *Sistema) Borrar()
