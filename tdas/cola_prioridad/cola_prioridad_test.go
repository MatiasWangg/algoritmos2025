package cola_prioridad_test

import (
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_MENSAJE_PANIC = "La cola esta vacia"
)

func compararNumeros(a, b int) int                { return a - b }
func compararCadenas(cadena1, cadena2 string) int { return strings.Compare(cadena1, cadena2) }
func TestHeapVacioComportamiento(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararCadenas), require.New(t)
	r.True(heap.EstaVacia(), "el heap deberia estar vacio inicialmente")
	r.EqualValues(0, heap.Cantidad(), "el heap deberia estar vacio inicialmente")
	r.PanicsWithValue(_MENSAJE_PANIC, func() { heap.VerMax() }, "deberia lanzar un panico al intentar ver el maximo de un heap vacio")
	r.PanicsWithValue(_MENSAJE_PANIC, func() { heap.Desencolar() }, "deberia lanzar un panico al intentar desencolar un elemento de un heap vacio")
}

func Test_NumerosYStrings(t *testing.T) {
	heapString, heapInt, r := TDAHeap.CrearHeap(compararCadenas), TDAHeap.CrearHeap(compararNumeros), require.New(t)
	heapString.Encolar("C")
	heapInt.Encolar(10)
	r.EqualValues(1, heapString.Cantidad(), "el heap de strings deberia contener un elemento despues de encolar 'C'")
	r.EqualValues(1, heapInt.Cantidad(), "el heap de enteros deberia contener un elemento despues de encolar 10")
	r.Equal("C", heapString.VerMax(), "el maximo del heap de strings deberia ser 'C'")
	r.Equal(10, heapInt.VerMax(), "el maximo del heap de enteros deberia ser 10")
}

func Test_ElementosNegativos(t *testing.T) {
	numerosConNegativos := []int{-10, -7, -4, -8, -12, -3, -1}
	heap, r := TDAHeap.CrearHeapArr(numerosConNegativos, compararNumeros), require.New(t)
	r.EqualValues(len(numerosConNegativos), heap.Cantidad(), "el heap de enteros deberia contener la misma cantidad de elementos que el arreglo")
	r.Equal(-1, heap.Desencolar(), "el primer elemento desencolado deberia ser -1")
	r.Equal(-3, heap.Desencolar(), "el proximo elemento desencolado deberia ser -3 (segundo maximo)")
}
func Test_ElementosDuplicados(t *testing.T) {
	numerosConDuplicados := []int{10, 7, 10, 8, 12, 3, 1, 12}
	heap, r := TDAHeap.CrearHeapArr(numerosConDuplicados, compararNumeros), require.New(t)
	r.EqualValues(len(numerosConDuplicados), heap.Cantidad(), "el heap de enteros deberia contener la misma cantidad de elementos que el arreglo")
	r.Equal(12, heap.Desencolar(), "el primer elemento desencolado deberia ser 12")
	r.Equal(12, heap.Desencolar(), "el segundo elemento desencolado tambien deberia ser 12")
}
func Test_HeapConStructs(t *testing.T) {
	type persona struct {
		nombre string
		edad   int
	}
	compararPersonas, personas := func(a, b persona) int { return a.edad - b.edad }, []persona{{"Juan", 30}, {"Ana", 25}, {"Pedro", 35}}
	heapPersonas, r := TDAHeap.CrearHeapArr(personas, compararPersonas), require.New(t)
	r.EqualValues(len(personas), heapPersonas.Cantidad(), "el heap de personas deberia contener la misma cantidad de elementos que el arreglo")
	r.Equal(persona{"Pedro", 35}, heapPersonas.Desencolar(), "la persona desencolada deberia ser la de mayor edad")
	r.Equal(persona{"Juan", 30}, heapPersonas.Desencolar(), "la persona desencolada deberia ser la de segunda mayor edad")
	r.Equal(persona{"Ana", 25}, heapPersonas.Desencolar(), "la persona desencolada deberia ser la de menor edad")

	r.True(heapPersonas.EstaVacia(), "el heap de personas deberia estar vacio despues de desencolar todos los elementos")
}
func Test_OrdenDeInsercion(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararNumeros), require.New(t)
	heap.Encolar(10)
	heap.Encolar(7)
	heap.Encolar(4)
	heap.Encolar(8)
	heap.Encolar(12)
	heap.Encolar(3)
	heap.Encolar(1)
	r.Equal(12, heap.Desencolar(), "el primer elemento desencolado deberia ser 12")
	r.Equal(10, heap.Desencolar(), "el segundo elemento desencolado deberia ser 10")
	r.Equal(8, heap.Desencolar(), "el tercer elemento desencolado deberia ser 8")
	r.Equal(7, heap.Desencolar(), "el cuarto elemento desencolado deberia ser 7")
	r.Equal(4, heap.Desencolar(), "el quinto elemento desencolado deberia ser 4")
	r.Equal(3, heap.Desencolar(), "el sexto elemento desencolado deberia ser 3")
	r.Equal(1, heap.Desencolar(), "el septimo elemento desencolado deberia ser 1")
}

func Test_Volumen(t *testing.T) {
	const cantidad = 10000
	heap, r := TDAHeap.CrearHeap(compararNumeros), require.New(t)

	for i := 0; i < cantidad; i++ {
		heap.Encolar(i)
	}
	r.EqualValues(cantidad, heap.Cantidad(), "el heap debería tener todos los elementos encolados")

	for i := cantidad - 1; i >= 0; i-- {
		r.EqualValues(i, heap.Desencolar(), "debería desencolar los elementos en orden descendente")
	}
	r.True(heap.EstaVacia(), "el heap debería quedar vacío tras desencolar todo")
}

func Test_HeapSortEnteros(t *testing.T) {
	arreglo := []int{5, 3, 8, 1, 9, 2}
	esperado := []int{1, 2, 3, 5, 8, 9}
	TDAHeap.HeapSort(arreglo, compararNumeros)
	require.Equal(t, esperado, arreglo, "HeapSort no ordenó correctamente")
}

func Test_HeapCreadoDesdeArregloVSManual(t *testing.T) {
	arreglo := []int{3, 1, 4, 2, 5}
	cmp := compararNumeros
	heapManual := TDAHeap.CrearHeap(cmp)
	for _, val := range arreglo {
		heapManual.Encolar(val)
	}
	heapArr := TDAHeap.CrearHeapArr(arreglo, cmp)

	for !heapManual.EstaVacia() {
		require.Equal(t, heapManual.Desencolar(), heapArr.Desencolar(), "Ambos heaps deberían desencolar los mismos elementos en el mismo orden")
	}
}
func Test_CrearHeapArrVacio(t *testing.T) {
	arregloVacio := []int{}
	heap := TDAHeap.CrearHeapArr(arregloVacio, compararNumeros)
	require.True(t, heap.EstaVacia(), "Heap creado desde arreglo vacío debería estar vacío")
	heap.Encolar(2)
	require.Equal(t, heap.Desencolar(), 2)
}

func Test_ReutilizarHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap(compararNumeros)
	heap.Encolar(1)
	heap.Encolar(2)
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())

	heap.Encolar(99)
	require.Equal(t, 99, heap.VerMax(), "El heap debería seguir funcionando después de vaciarse")
}
func Test_VolumenVerMax(t *testing.T) {
	const N = 5000
	heap := TDAHeap.CrearHeap(compararNumeros)
	for i := 0; i < N; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax(), "El máximo debería ser el último insertado")
	}
}
func Test_HeapSortDuplicados(t *testing.T) {
	arreglo := []int{4, 4, 4, 4}
	TDAHeap.HeapSort(arreglo, compararNumeros)
	require.Equal(t, []int{4, 4, 4, 4}, arreglo)
}

func Test_HeapSortConStruct(t *testing.T) {
	type persona struct {
		nombre string
		edad   int
	}
	personas := []persona{{"Ana", 25}, {"Juan", 30}, {"Pedro", 20}}
	cmp := func(a, b persona) int { return a.edad - b.edad }

	TDAHeap.HeapSort(personas, cmp)

	require.Equal(t, []persona{{"Pedro", 20}, {"Ana", 25}, {"Juan", 30}}, personas)
}
