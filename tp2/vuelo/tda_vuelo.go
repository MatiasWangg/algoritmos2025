package vuelo

import (
	"time"
)

type Vuelo struct {
    Codigo           string 
    Aerolinea        string 
    Origen           string 
    Destino          string 
    Matricula        string 
    Prioridad        int   
    Fecha            time.Time 
    RetrasoSalida    int    
    TiempoVuelo      int    
    Cancelado        bool   
}

func CrearVuelo(codigo, aerolinea, origen, destino, matricula string, prioridad, retrasoSalida, tiempoVuelo int, fecha time.Time, cancelado bool) *Vuelo {
	vuelo := new(Vuelo)
	vuelo.Codigo = codigo
	vuelo.Aerolinea = aerolinea
	vuelo.Origen = origen
	vuelo.Destino = destino
	vuelo.Matricula = matricula
	vuelo.Prioridad = prioridad
	vuelo.Fecha = fecha
	vuelo.RetrasoSalida = retrasoSalida
	vuelo.TiempoVuelo = tiempoVuelo
	vuelo.Cancelado = cancelado
	return vuelo
}

func (v *Vuelo) EstaCanceladoInt() int {
	if v.Cancelado {
		return 1
	}
	return 0
}

func (v *Vuelo) ObtenerOrigen() string {
	return v.Origen
}

func (v *Vuelo) ObtenerDestino() string {
	return v.Destino
}

func (v *Vuelo) ObtenerFecha() time.Time {
	return v.Fecha
}

