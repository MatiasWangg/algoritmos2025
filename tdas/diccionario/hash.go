package diccionario

import (
	"fmt"
)

type estadoParaClaveDato int

const (
	VACIO = estadoParaClaveDato(iota)
	BORRADO
	OCUPADO
	_TAMANIO_INICIAL        = 100
	_FACTOR_CARGA           = 0.7
	_FACTOR_ESCALA          = 2
	_MENSAJE_PANIC_HASH     = "La clave no pertenece al diccionario"
	_MENSAJE_PANIC_ITERADOR = "El iterador termino de iterar"
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoParaClaveDato
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}
type iteradorHash[K comparable, V any] struct {
	hash   *hashCerrado[K, V]
	indice int
}

func crearTablaHash[K comparable, V any](tamanio int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], tamanio)
}
func (h *hashCerrado[K, V]) gestionarCapacidad() {
	tamanioTabla := h.tam
	if carga := float64(h.cantidad+h.borrados) / float64(h.tam); carga > _FACTOR_CARGA {
		tamanioTabla *= _FACTOR_ESCALA
	} else if carga < 1-_FACTOR_CARGA && tamanioTabla > _TAMANIO_INICIAL {
		tamanioTabla /= _FACTOR_ESCALA
	} else {
		return
	}

	nuevaTabla := crearTablaHash[K, V](tamanioTabla)
	for _, celda := range h.tabla {
		if celda.estado == OCUPADO {
			indice := funcionHash(convertirABytes(celda.clave), tamanioTabla)
			for nuevaTabla[indice].estado == OCUPADO {
				indice = (indice + 1) % uint64(tamanioTabla)
			}
			nuevaTabla[indice] = celda
		}
	}
	h.tabla = nuevaTabla
	h.tam = tamanioTabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashCerrado[K, V]{
		tabla:    crearTablaHash[K, V](_TAMANIO_INICIAL),
		tam:      _TAMANIO_INICIAL,
		cantidad: 0,
		borrados: 0,
	}
	return hash
}
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
func funcionHash(key []byte, tam int) uint64 {
	var hash uint64 = 2166136261
	for _, c := range key {
		hash ^= uint64(c)
		hash *= 16777619
	}
	return hash % uint64(tam)
}

func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, existe := h.buscarYVerificar(clave)
	return existe
}

func (h *hashCerrado[K, V]) Obtener(clave K) V {
	if indice, existe := h.buscarYVerificar(clave); existe {
		return h.tabla[indice].dato
	}
	panic(_MENSAJE_PANIC_HASH)
}

func (h *hashCerrado[K, V]) Guardar(clave K, dato V) {
	indice, existe := h.buscarYVerificar(clave)
	if !existe {
		h.tabla[indice].estado = OCUPADO
		h.tabla[indice].clave = clave
		h.cantidad++
	}
	h.tabla[indice].dato = dato
	h.gestionarCapacidad()
}

func (h *hashCerrado[K, V]) Borrar(clave K) V {
	if indice, existe := h.buscarYVerificar(clave); existe {
		datoB := h.tabla[indice].dato
		h.tabla[indice].estado = BORRADO
		h.cantidad--
		h.borrados++
		h.gestionarCapacidad()
		return datoB
	}
	panic(_MENSAJE_PANIC_HASH)
}

func (h *hashCerrado[K, V]) Cantidad() int { return h.cantidad }

func (h *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range h.tabla {
		if celda.estado == OCUPADO && !visitar(celda.clave, celda.dato) {
			break
		}
	}
}
func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHash[K, V]{
		hash:   hash,
		indice: buscarIndiceOcupado(hash.tabla, 0),
	}
}

func (iter *iteradorHash[K, V]) HaySiguiente() bool {
	indice := buscarIndiceOcupado(iter.hash.tabla, iter.indice)
	return indice < len(iter.hash.tabla) && iter.hash.tabla[indice].estado == OCUPADO
}

func (iter *iteradorHash[K, V]) VerActual() (K, V) {
	if iter.HaySiguiente() {
		return iter.hash.tabla[iter.indice].clave, iter.hash.tabla[iter.indice].dato
	}
	panic(_MENSAJE_PANIC_ITERADOR)
}
func (iter *iteradorHash[K, V]) Siguiente() {
	if iter.HaySiguiente() {
		iter.indice = buscarIndiceOcupado(iter.hash.tabla, iter.indice+1)
		return
	}
	panic(_MENSAJE_PANIC_ITERADOR)
}

func (hash *hashCerrado[K, V]) buscar(clave K) int {
	indice := funcionHash(convertirABytes(clave), hash.tam)
	for {
		if hash.tabla[indice].estado == OCUPADO && hash.tabla[indice].clave == clave ||
			hash.tabla[indice].estado == VACIO {
			return int(indice)
		}
		indice = (indice + 1) % uint64(hash.tam)
	}
}

func buscarIndiceOcupado[K comparable, V any](tabla []celdaHash[K, V], indiceActual int) int {
	for i := indiceActual; i < len(tabla); i++ {
		if tabla[i].estado == OCUPADO {
			return i
		}
	}
	return indiceActual
}
func (h *hashCerrado[K, V]) buscarYVerificar(clave K) (int, bool) {
	indice := h.buscar(clave)
	existe := h.cantidad != 0 && h.tabla[indice].clave == clave
	return indice, existe
}
