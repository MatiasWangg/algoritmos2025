package lista

type nodoLista[T any] struct {
	dato T
	sig  *nodoLista[T]
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	nodoLista := new(nodoLista[T])

	nodoLista.dato = dato
	nodoLista.sig = nil

	return nodoLista
}

type lista_enlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(lista_enlazada[T])

	return lista
}

func (lista *lista_enlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil && lista.largo == 0
}

func (lista *lista_enlazada[T]) InsertarPrimero(dato T) {
	nodo := nodoCrear(dato)
	if lista.EstaVacia() {
		lista.ultimo = nodo
	}
	nodo.sig = lista.primero
	lista.primero = nodo

	lista.largo++
}

func (lista *lista_enlazada[T]) InsertarUltimo(dato T) {
	nodo := nodoCrear(dato)
	if lista.EstaVacia() {
		lista.primero = nodo
	} else {
		lista.ultimo.sig = nodo
	}
	lista.ultimo = nodo
	lista.largo++
}

func (lista *lista_enlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	borrado := lista.primero
	lista.primero = lista.primero.sig
	lista.largo--
	if lista.largo == 0 {
		lista.ultimo = nil
	}
	return borrado.dato
}

func (lista *lista_enlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *lista_enlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *lista_enlazada[T]) Largo() int {
	return lista.largo
}

// Iterador Interno

func (lista *lista_enlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil && visitar(actual.dato) {
		actual = actual.sig
	}
}

//Estructura y Creador del Iterador externo

func (lista *lista_enlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iterListaEnlazda[T])
	iter.actual = lista.primero
	iter.anterior = nil
	iter.lista = lista
	return iter
}

type iterListaEnlazda[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *lista_enlazada[T]
}

func (iter *iterListaEnlazda[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazda[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazda[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.sig
}

func (iter *iterListaEnlazda[T]) Insertar(elem T) {
	nodo := nodoCrear(elem)

	if iter.actual == nil {
		iter.lista.ultimo = nodo
	}
	nodo.sig = iter.actual

	if iter.actual == iter.lista.primero {
		iter.lista.primero = nodo
	} else {
		iter.anterior.sig = nodo
	}

	iter.actual = nodo
	iter.lista.largo++
}

func (iter *iterListaEnlazda[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	borrado := iter.actual.dato

	if iter.actual == iter.lista.ultimo {
		iter.lista.ultimo = iter.anterior
	}

	if iter.actual == iter.lista.primero {
		iter.lista.primero = iter.actual.sig
	} else {
		iter.anterior.sig = iter.actual.sig
	}

	iter.actual = iter.actual.sig
	iter.lista.largo--

	return borrado
}
