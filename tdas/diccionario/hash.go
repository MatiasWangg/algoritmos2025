package diccionario

import (
	"fmt"
)

type estadoParaClaveDato int

const (
	_VACIO = estadoParaClaveDato(iota)
	_BORRADO
	_OCUPADO

	_TAMANIO_INICIAL        = 100
	_FACTOR_CARGA_MAX       = 0.7
	_FACTOR_CARGA_MIN       = 0.1
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

func (h *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	tablaAnterior := h.tabla
	h.tabla = crearTablaHash[K, V](nuevoTam)
	h.tam = nuevoTam
	h.cantidad = 0
	h.borrados = 0

	for _, celda := range tablaAnterior {
		if celda.estado == _OCUPADO {
			h.Guardar(celda.clave, celda.dato)
		}
	}
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{
		tabla:    crearTablaHash[K, V](_TAMANIO_INICIAL),
		tam:      _TAMANIO_INICIAL,
		cantidad: 0,
		borrados: 0,
	}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// funcionHash FNV-1a link de docu (http://www.isthe.com/chongo/tech/comp/fnv/#FNV-1)
func funcionHash(key []byte, tam int) int {
	var hash uint64 = 2166136261
	for _, c := range key {
		hash ^= uint64(c)
		hash *= 16777619
	}
	return int(hash % uint64(tam))
}

func (h *hashCerrado[K, V]) buscar(clave K) int {
	indice := funcionHash(convertirABytes(clave), h.tam)

	for {
		if h.tabla[indice].estado == _OCUPADO && h.tabla[indice].clave == clave ||
			h.tabla[indice].estado == _VACIO {
			return indice
		}
		indice = (indice + 1) % h.tam
	}
}

func (h *hashCerrado[K, V]) buscarYVerificar(clave K) (int, bool) {
	indice := h.buscar(clave)
	existe := h.cantidad != 0 && h.tabla[indice].estado == _OCUPADO && h.tabla[indice].clave == clave
	return indice, existe
}

func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, existe := h.buscarYVerificar(clave)
	return existe
}

func (h *hashCerrado[K, V]) Obtener(clave K) V {
	indice, existe := h.buscarYVerificar(clave)
	if !existe {
		panic(_MENSAJE_PANIC_HASH)
	}
	return h.tabla[indice].dato
}

func (h *hashCerrado[K, V]) Guardar(clave K, dato V) {
	indice, existe := h.buscarYVerificar(clave)
	if !existe {
		h.tabla[indice].estado = _OCUPADO
		h.tabla[indice].clave = clave
		h.cantidad++
	}
	h.tabla[indice].dato = dato
	carga := float64(h.cantidad+h.borrados) / float64(h.tam)
	if carga > _FACTOR_CARGA_MAX {
		h.redimensionar(h.tam * _FACTOR_ESCALA)
	}
}

func (h *hashCerrado[K, V]) Borrar(clave K) V {
	indice, existe := h.buscarYVerificar(clave)
	if !existe {
		panic(_MENSAJE_PANIC_HASH)
	}
	datoBorrado := h.tabla[indice].dato
	h.tabla[indice].estado = _BORRADO
	h.cantidad--
	h.borrados++

	carga := float64(h.cantidad+h.borrados) / float64(h.tam)
	if carga < _FACTOR_CARGA_MIN && h.tam > _TAMANIO_INICIAL {
		h.redimensionar(h.tam / _FACTOR_ESCALA)
	}
	return datoBorrado
}

func (h *hashCerrado[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range h.tabla {
		if celda.estado == _OCUPADO {
			if !visitar(celda.clave, celda.dato) {
				break
			}
		}
	}
}
func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHash[K, V]{
		hash:   h,
		indice: buscarIndiceOcupado(h.tabla, 0),
	}
}

func buscarIndiceOcupado[K comparable, V any](tabla []celdaHash[K, V], indiceActual int) int {
	for i := indiceActual; i < len(tabla); i++ {
		if tabla[i].estado == _OCUPADO {
			return i
		}
	}
	return len(tabla)
}

func (iter *iteradorHash[K, V]) HaySiguiente() bool {
	return iter.indice < len(iter.hash.tabla) &&
		iter.hash.tabla[iter.indice].estado == _OCUPADO
}

func (iter *iteradorHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITERADOR)
	}
	return iter.hash.tabla[iter.indice].clave, iter.hash.tabla[iter.indice].dato
}

func (iter *iteradorHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_PANIC_ITERADOR)
	}
	iter.indice = buscarIndiceOcupado(iter.hash.tabla, iter.indice+1)
}
