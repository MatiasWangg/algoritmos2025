package diccionario

import (
	p "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.raiz = nil
	abb.cmp = funcion_cmp
	return abb
}

func (abb *abb[K, V]) buscar(clave K, nodo *nodoAbb[K, V], padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nodo, padre
	}
	comparacion := abb.cmp(clave, nodo.clave)

	if comparacion == 0 {
		return nodo, padre
	} else if comparacion < 0 {
		return abb.buscar(clave, nodo.izquierdo, nodo)
	} else {
		return abb.buscar(clave, nodo.derecho, nodo)
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.buscar(clave, abb.raiz, nil)
	return nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := abb.buscar(clave, abb.raiz, nil)

	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	return nodo.dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo, padre := abb.buscar(clave, abb.raiz, nil)
	if nodo != nil {
		nodo.dato = dato
		return
	}
	nuevoNodo := crearNodoABB[K, V](clave, dato)
	if padre == nil {
		abb.raiz = nuevoNodo
	} else {
		if abb.cmp(clave, padre.clave) < 0 {
			padre.izquierdo = nuevoNodo
		} else {
			padre.derecho = nuevoNodo
		}
	}
	abb.cantidad++
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, padre := abb.buscar(clave, abb.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	dato := nodo.dato

	if nodo.izquierdo == nil && nodo.derecho == nil {
		if padre != nil {
			if abb.cmp(clave, padre.clave) < 0 {
				padre.izquierdo = nil
			} else {
				padre.derecho = nil
			}
		} else {
			abb.raiz = nil
		}
	} else if nodo.izquierdo == nil || nodo.derecho == nil {
		var hijo *nodoAbb[K, V]
		if nodo.izquierdo != nil {
			hijo = nodo.izquierdo
		} else {
			hijo = nodo.derecho
		}

		if padre != nil {
			if abb.cmp(clave, padre.clave) < 0 {
				padre.izquierdo = hijo
			} else {
				padre.derecho = hijo
			}
		} else {
			abb.raiz = hijo
		}
	} else {
		padreSustituto := nodo
		sustituto := nodo.derecho

		for sustituto.izquierdo != nil {
			padreSustituto = sustituto
			sustituto = sustituto.izquierdo
		}

		nodo.clave = sustituto.clave
		nodo.dato = sustituto.dato

		if padreSustituto.izquierdo == sustituto {
			padreSustituto.izquierdo = sustituto.derecho
		} else {
			padreSustituto.derecho = sustituto.derecho
		}
	}

	abb.cantidad--
	return dato
}

//----------------------------------Iteradores-----------------------------------------

//interno

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

//Interno por rangos

func (nodo *nodoAbb[K, V]) iterarRangos(desde *K, hasta *K, cmp func(K, K) int, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	cmpDesde := 0
	if desde != nil {
		cmpDesde = cmp(nodo.clave, *desde)
	}

	cmpHasta := 0
	if hasta != nil {
		cmpHasta = cmp(nodo.clave, *hasta)
	}

	if cmpDesde >= 0 {
		if !nodo.izquierdo.iterarRangos(desde, hasta, cmp, visitar) {
			return false
		}
	}
	if cmpDesde >= 0 && cmpHasta <= 0 {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if cmpHasta <= 0 {
		if !nodo.derecho.iterarRangos(desde, hasta, cmp, visitar) {
			return false
		}
	}

	return true
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarRangos(desde, hasta, abb.cmp, visitar)
}

// Externo---------------------------------------------------------------------------------------
func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

// Externo por rangos---------------------------------------------------------------------------------------

type iteradorArbol[K comparable, V any] struct {
	arbol *abb[K, V]
	pila  p.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(K, K) int
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	abbiter := new(iteradorArbol[K, V])
	abbiter.pila = p.CrearPilaDinamica[*nodoAbb[K, V]]()
	abbiter.desde = desde
	abbiter.hasta = hasta
	abbiter.arbol = abb
	abbiter.cmp = abb.cmp
	abbiter.apilarDesde(abb.raiz)
	return abbiter
}

func (abbiter *iteradorArbol[K, V]) HaySiguiente() bool {
	return !abbiter.pila.EstaVacia()
}

func (abbiter *iteradorArbol[K, V]) VerActual() (K, V) {
	if !abbiter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := abbiter.pila.VerTope()
	return actual.clave, actual.dato
}

func (abbiter *iteradorArbol[K, V]) Siguiente() {
	if !abbiter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := abbiter.pila.Desapilar()
	if nodo.derecho != nil {
		abbiter.apilarDesde(nodo.derecho)
	}
}

func (abbiter *iteradorArbol[K, V]) apilarDesde(nodo *nodoAbb[K, V]) {

	for nodo != nil {
		if abbiter.desde != nil && abbiter.cmp(nodo.clave, *abbiter.desde) < 0 {
			nodo = nodo.derecho
		} else {
			if abbiter.hasta == nil || abbiter.cmp(nodo.clave, *abbiter.hasta) <= 0 {
				abbiter.pila.Apilar(nodo)
			}
			nodo = nodo.izquierdo
		}
	}
}
