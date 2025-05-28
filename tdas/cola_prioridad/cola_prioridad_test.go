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
func TestHeapVacio(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararCadenas), require.New(t)
	r.EqualValues(0, heap.Cantidad(), "el heap deberia estar vacio inicialmente")
	r.PanicsWithValue(_MENSAJE_PANIC, func() { heap.VerMax() }, "deberia lanzar un panico al intentar ver el maximo de un heap vacio")
	r.PanicsWithValue(_MENSAJE_PANIC, func() { heap.Desencolar() }, "deberia lanzar un panico al intentar desencolar un elemento de un heap vacio")
}
func Test_EstaVacia(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararCadenas), require.New(t)
	r.True(heap.EstaVacia(), "el heap deberia estar vacio inicialmente")
	heap.Encolar("A")
	r.False(heap.EstaVacia(), "el heap no deberia estar vacio despues de encolar un elemento")
	heap.Desencolar()
	r.True(heap.EstaVacia(), "el heap deberia estar vacio despues de desencolar el unico elemento")
}
func Test_Cantidad(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararCadenas), require.New(t)
	r.EqualValues(0, heap.Cantidad(), "el heap deberia tener cero elementos inicialmente")
	heap.Encolar("A")
	r.EqualValues(1, heap.Cantidad(), "el heap deberia tener un elemento despues de encolar uno")
	heap.Encolar("B")
	r.EqualValues(2, heap.Cantidad(), "el heap deberia tener dos elementos despues de encolar otro")
}
func Test_VerMax(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararCadenas), require.New(t)

	heap.Encolar("C")
	heap.Encolar("A")
	heap.Encolar("B")

	r.Equal("C", heap.VerMax(), "el maximo del heap deberia ser 'C'")
	r.EqualValues(3, heap.Cantidad(), "el heap deberia contener tres elementos despues de encolarlos")
}
func Test_EncolarNumeroYString(t *testing.T) {
	heapString, heapInt, r := TDAHeap.CrearHeap(compararCadenas), TDAHeap.CrearHeap(compararNumeros), require.New(t)
	heapString.Encolar("C")
	heapInt.Encolar(10)
	r.EqualValues(1, heapString.Cantidad(), "el heap de strings deberia contener un elemento despues de encolar 'C'")
	r.EqualValues(1, heapInt.Cantidad(), "el heap de enteros deberia contener un elemento despues de encolar 10")
	r.Equal("C", heapString.VerMax(), "el maximo del heap de strings deberia ser 'C'")
	r.Equal(10, heapInt.VerMax(), "el maximo del heap de enteros deberia ser 10")
}
func Test_EncolarYDesencolar(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararCadenas), require.New(t)
	heap.Encolar("C")
	heap.Encolar("A")
	heap.Encolar("B")
	r.EqualValues(3, heap.Cantidad(), "el heap deberia contener tres elementos despues de encolarlos")
	r.Equal("C", heap.VerMax(), "el maximo del heap deberia ser 'C'")
	r.Equal("C", heap.Desencolar(), "el maximo del heap deberia ser 'C' despues de desencolarlo")
	r.EqualValues(2, heap.Cantidad(), "el heap deberia contener dos elementos despues de desencolar el maximo")
	r.Equal("B", heap.VerMax(), "el maximo del heap deberia ser 'B'")
	r.Equal("B", heap.Desencolar(), "el maximo del heap deberia ser 'B' despues de desencolarlo")
	r.EqualValues(1, heap.Cantidad(), "el heap deberia contener un elemento despues de desencolar el maximo")
	r.Equal("A", heap.VerMax(), "el maximo del heap deberia ser 'A'")
	r.Equal("A", heap.Desencolar(), "el maximo del heap deberia ser 'A' despues de desencolarlo")
	r.EqualValues(0, heap.Cantidad(), "el heap deberia estar vacio despues de desencolar el maximo")
}
func Test_EncolarYDesencolarInt(t *testing.T) {
	heap, r := TDAHeap.CrearHeap(compararNumeros), require.New(t)
	heap.Encolar(3)
	heap.Encolar(1)
	heap.Encolar(2)
	r.EqualValues(3, heap.Cantidad(), "el heap deberia contener tres elementos despues de encolarlos")
	r.Equal(3, heap.VerMax(), "el maximo del heap deberia ser 3")
	r.Equal(3, heap.Desencolar(), "el maximo del heap deberia ser 3 despues de desencolarlo")
	r.EqualValues(2, heap.Cantidad(), "el heap deberia contener dos elementos despues de desencolar el maximo")
	r.Equal(2, heap.VerMax(), "el maximo del heap deberia ser 2")
	r.Equal(2, heap.Desencolar(), "el maximo del heap deberia ser 2 despues de desencolarlo")
	r.EqualValues(1, heap.Cantidad(), "el heap deberia contener un elemento despues de desencolar el maximo")
	r.Equal(1, heap.VerMax(), "el maximo del heap deberia ser 1")
	r.Equal(1, heap.Desencolar(), "el maximo del heap deberia ser 1 despues de desencolarlo")
	r.EqualValues(0, heap.Cantidad(), "el heap deberia estar vacio despues de desencolar el maximo")
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

func Test_CrearHeapArrNumeros(t *testing.T) {
	numeros := []int{10, 7, 4, 8, 12, 3, 1}
	numerosOrdenado := []int{12, 10, 8, 7, 4, 3, 1}
	heap, r := TDAHeap.CrearHeapArr(numeros, compararNumeros), require.New(t)
	r.EqualValues(len(numeros), heap.Cantidad(), "el heap de enteros deberia contener la misma cantidad de elementos que el arreglo")
	for _, v := range numerosOrdenado {
		r.Equal(v, heap.Desencolar(), "el elemento desencolado deberia ser igual al elemento en la misma posicion del arreglo ordenado")
	}
	r.True(heap.EstaVacia(), "el heap de enteros deberia estar vacio despues de desencolar todos los elementos")
}

/* */
var _ARREGLO_TEST = []int{7, 1, 2, 3, 30, 10, 80}

func TestHeapVacia(t *testing.T) {
	heap, require := TDAHeap.CrearHeap(compararNumeros), require.New(t)
	require.True(heap.EstaVacia(), "La cola deberia estar vacia ni bien es creada.")
	require.EqualValues(heap.Cantidad(), 0, "La cola deberia tener cantidad 0 ni bien es creada.")
}
func TestHeapArr(t *testing.T) {
	heap, require := TDAHeap.CrearHeapArr(_ARREGLO_TEST, compararNumeros), require.New(t)
	require.False(heap.EstaVacia(), "La cola deberia estar vacia ni bien es creada.")
	heap.Encolar(555)
}
func TestHeapMax(t *testing.T) {
	heap, require := TDAHeap.CrearHeapArr(_ARREGLO_TEST, compararNumeros), require.New(t)
	heap.Encolar(555)
	require.EqualValues(555, heap.VerMax())
}
func Test_HeapMaxEncolandoElementoAElemento(t *testing.T) {
	heap, require := TDAHeap.CrearHeap(compararNumeros), require.New(t)
	for _, v := range _ARREGLO_TEST {
		heap.Encolar(v)
	}
	heap.Desencolar()
	heap.Encolar(555)
	require.EqualValues(555, heap.VerMax())
}
