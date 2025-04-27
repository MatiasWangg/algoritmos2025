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