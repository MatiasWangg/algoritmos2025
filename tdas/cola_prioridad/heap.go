package cola_prioridad

const (
	_CAPACIDAD_MINIMA    = 10
	_MENSAJE_PANIC       = "La cola esta vacia"
	_FACTOR_CAMBIO       = 2
	_MODIFICADOR_TAMANIO = 4
	_FACTOR_HIJO         = 2
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, _CAPACIDAD_MINIMA),
		cant:  0,
		cmp:   funcion_cmp,
	}
}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	capacidad := len(arreglo)
	if capacidad < _CAPACIDAD_MINIMA {
		capacidad = _CAPACIDAD_MINIMA
	}

	datos := make([]T, capacidad)
	copy(datos, arreglo)

	heapArr := &colaConPrioridad[T]{
		datos: datos,
		cant:  len(arreglo),
		cmp:   funcion_cmp,
	}

	heapify(heapArr.datos[:heapArr.cant], heapArr.cmp)
	return heapArr
}

func (heap *colaConPrioridad[T]) EstaVacia() bool { return heap.cant == 0 }
func (heap *colaConPrioridad[T]) Encolar(dato T) {
	if heap.cant == len(heap.datos) {
		tamanioHeap := len(heap.datos)
		tamanioHeap *= _FACTOR_CAMBIO
		heap.redimensionarCola(tamanioHeap)
	}
	heap.datos[heap.cant], heap.cant = dato, heap.cant+1
	heap.upHeap(heap.cant - 1)
}

func (heap *colaConPrioridad[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	return heap.datos[0]
}
func (heap *colaConPrioridad[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(_MENSAJE_PANIC)
	}
	maximo := heap.VerMax()
	swap(heap.datos, 0, heap.cant-1)
	heap.cant--
	downheap(heap.datos, 0, heap.cant, heap.cmp)
	if len(heap.datos)/_MODIFICADOR_TAMANIO >= heap.cant && len(heap.datos) > _CAPACIDAD_MINIMA {
		tamanioHeap := len(heap.datos)
		tamanioHeap /= _FACTOR_CAMBIO
		heap.redimensionarCola(tamanioHeap)
	}
	return maximo
}
func (heap *colaConPrioridad[T]) Cantidad() int { return heap.cant }

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	// Convertir el arreglo en un heap máximo
	heapify(elementos, funcion_cmp)

	// Ordenar el arreglo in-place
	for i := len(elementos) - 1; i > 0; i-- {
		swap(elementos, 0, i)
		downheap(elementos, 0, i, funcion_cmp)
	}
}

func swap[T any](arr []T, i, j int) { arr[i], arr[j] = arr[j], arr[i] }

func heapify[T any](datos []T, cmp func(T, T) int) {
	n := len(datos)
	for i := (n / _FACTOR_HIJO) - 1; i >= 0; i-- {
		downheap(datos, i, n, cmp)
	}
}
func downheap[T any](datos []T, i, n int, cmp func(T, T) int) {
	maxIndice := i
	indHijoIzq := 2*i + 1
	indHijoDer := 2*i + 2

	if indHijoIzq < n && cmp(datos[indHijoIzq], datos[maxIndice]) > 0 {
		maxIndice = indHijoIzq
	}
	if indHijoDer < n && cmp(datos[indHijoDer], datos[maxIndice]) > 0 {
		maxIndice = indHijoDer
	}
	if maxIndice != i {
		swap(datos, i, maxIndice)
		downheap(datos, maxIndice, n, cmp)
	}
}
func (heap *colaConPrioridad[T]) upHeap(i int) {
	for i != 0 {
		padre := (i - 1) / 2
		if heap.cmp(heap.datos[i], heap.datos[padre]) <= 0 {
			break
		}
		swap(heap.datos, i, padre)
		i = padre
	}
}

func (heap *colaConPrioridad[T]) redimensionarCola(nuevoTamaño int) {
	nuevosDatos := make([]T, nuevoTamaño)
	copy(nuevosDatos, heap.datos[:heap.cant])
	heap.datos = nuevosDatos
}
