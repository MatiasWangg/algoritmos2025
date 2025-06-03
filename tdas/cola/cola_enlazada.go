package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodoCola[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil

	return cola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return ((cola.primero == nil) && (cola.ultimo == nil))
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	primero := cola.primero.dato
	return primero

}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nodo := crearNodoCola(elemento)
	if cola.EstaVacia() {
		cola.primero = nodo
	} else {
		cola.ultimo.prox = nodo
	}
	cola.ultimo = nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	desencolado := cola.primero.dato
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return desencolado
}
