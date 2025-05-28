package pila

const (
	CAPACIDAD_INICIAL int = 5
	AGRANDAR          int = 2
	REDUCIR           int = 2
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_INICIAL)
	pila.cantidad = 0

	return pila
}

func (pila *pilaDinamica[T]) redimensionar(nueva_capacidad int) {
	nuevo := make([]T, nueva_capacidad)
	copy(nuevo, pila.datos)
	pila.datos = nuevo
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	tope := pila.datos[pila.cantidad-1]
	return tope
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) * AGRANDAR)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	if pila.cantidad*4 <= cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) / REDUCIR)
	}

	desapilado := pila.datos[pila.cantidad-1]
	pila.cantidad--
	return desapilado
}
