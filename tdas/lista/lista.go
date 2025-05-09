package lista

type Lista[T any] interface {

	//EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario
	EstaVacia() bool

	//InsertarPrimero inserta el elemento al inicio de la lista
	InsertarPrimero(T)

	//InsertarUltimo inserta el elemento al final de la lista
	InsertarUltimo(T)

	//BorrarPrimero borra el primer elemento de la lista. Si la lista esta vacia, entra en panico con un mensaje "La lista esta vacia"
	BorrarPrimero() T

	//VerPrimero devuelve el primer elemento de la lista. Si la lista esta vacia, entra en panico con un mensaje "La lista esta vacia"
	VerPrimero() T

	//VerUltimo devuelve el ultimo elemento de la lista. Si la lista esta vacia, entra en panico con un mensaje "La lista esta vacia"
	VerUltimo() T

	//Largo devulve la cantidad de elementos de la lista
	Largo() int

	//Iterar dicha funcion se aplica a cada elementos de la lista, hasta que se termine o la funcion visitar devuelva false
	Iterar(visitar func(T) bool)

	//Devuelve una instancia de IteradorLista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//VerActual devuelve el elemento actual de la iteracion. Si se la invoca sobre un iterador que ya haya iterado todos los elementos, entra en panico con un mensaje "El iterador termino de iterar"
	VerActual() T

	//HaySiguiente devulve true si hay algun elemento para observar en la posicion
	HaySiguiente() bool

	//Siguiente avanza una posicion en la iteracion. Si se la invoca sobre un iterador que ya haya iterado todos los elementos, entra en panico con un mensaje "El iterador termino de iterar"
	Siguiente()

	//Insertar inserta un elemento en la posicion actual y justo antes del elemento apuntado actualmente
	Insertar(T)

	//Borrar borra el elemento de la posicion actual y lo devuelve. Luego de borrar, el iterador apunta al siguiente elemento de la lista. Si se la invoca sobre un iterador que ya haya iterado todos los elementos, entra en panico con un mensaje "El iterador termino de iterar"
	Borrar() T
}
