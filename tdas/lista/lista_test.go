package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLista_NuevaListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "Deberia estar vacia al ser creada")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "BorrarPrimero deberia causar panico en una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "VerPrimero deberia causar panico en una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "VerUltimo deberia causar panico en una lista vacia")
}

func TestLista_InsertarBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(0)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia despues de insertar")
	require.Equal(t, 0, lista.VerPrimero(), "El primer valor debería ser el insertado")

	lista.InsertarUltimo(1)
	require.Equal(t, 0, lista.VerPrimero(), "El primer valor deberia seguir siendo el mismo")

	borrado := lista.BorrarPrimero()
	require.Equal(t, 0, borrado, "Se deberia haber borrado el primero")
	require.Equal(t, 1, lista.VerPrimero(), "El primer valor ahora deberia ser el segundo insertado")

	borrado = lista.BorrarPrimero()
	require.Equal(t, 1, borrado, "Se deberia haber borrado el ultimo")
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todo")
}

func TestLista_IntercaladoYCasosBorde(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarUltimo("uno")
	require.Equal(t, "uno", lista.VerPrimero(), "un unico elemento insertado al final deberia ser el primero")

	require.Equal(t, "uno", lista.BorrarPrimero(), "debe borrarse el unico elemento")
	require.True(t, lista.EstaVacia(), "lista deberia estar vacia tras borrar el unico elemento")

	lista.InsertarUltimo("dos")
	lista.InsertarPrimero("uno")
	require.Equal(t, "uno", lista.VerPrimero(), "insertarprimero deberia ubicar el nuevo primero")
	require.Equal(t, "uno", lista.BorrarPrimero(), "deberia borrarse el primero actual")

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia(), "lista deberia estar vacia nuevamente")

	lista.InsertarPrimero("final")
	require.False(t, lista.EstaVacia(), "despues de insertar, la lista no deberia estar vacia")
	require.Equal(t, "final", lista.VerPrimero(), "el unico elemento deberia ser final")
}

func TestLista_VolumenInsertarUltimo(t *testing.T) {
	tam := 1000
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, 0, lista.VerPrimero(), "el primero deberia ser 0")
		require.Equal(t, i, lista.VerUltimo(), "el ultimo deberia ser el ultimo insertado")
		require.False(t, lista.EstaVacia(), "la lista no deberia estar vacia")
	}

	for j := 0; j < tam; j++ {
		require.Equal(t, j, lista.VerPrimero(), "el primero deberia ser el actual a borrar")
		require.Equal(t, tam-1, lista.VerUltimo(), "el ultimo deberia seguir siendo el mismo")
		require.False(t, lista.EstaVacia(), "la lista no deberia estar vacia")
		require.Equal(t, j, lista.BorrarPrimero(), "el elemento borrado deberia coincidir con el esperado")
	}

	require.True(t, lista.EstaVacia(), "la lista deberia estar vacia al final")
}

func TestLista_VolumenInsertarPrimero(t *testing.T) {
	tam := 1000
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < tam; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero(), "el primero deberia ser el ultimo insertado")
		require.Equal(t, 0, lista.VerUltimo(), "el ultimo deberia seguir siendo el primero insertado")
		require.False(t, lista.EstaVacia(), "la lista no deberia estar vacia")
	}

	for j := tam - 1; j >= 0; j-- {
		require.Equal(t, j, lista.VerPrimero(), "el primero deberia ser el actual a borrar")
		require.Equal(t, 0, lista.VerUltimo(), "el ultimo deberia seguir siendo el mismo")
		require.False(t, lista.EstaVacia(), "la lista no deberia estar vacia")
		require.Equal(t, j, lista.BorrarPrimero(), "el elemento borrado deberia coincidir con el esperado")
	}

	require.True(t, lista.EstaVacia(), "la lista deberia estar vacia al final")
}

func TestLista_Generica(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaInt.InsertarUltimo(10)
	require.Equal(t, 10, listaInt.BorrarPrimero(), "Deberia borrar el entero 10")

	listaString := TDALista.CrearListaEnlazada[string]()
	listaString.InsertarUltimo("Hola")
	require.Equal(t, "Hola", listaString.BorrarPrimero(), "Deberia borrar la cadena 'Hola'")

	listaBool := TDALista.CrearListaEnlazada[bool]()
	listaBool.InsertarUltimo(true)
	require.Equal(t, true, listaBool.BorrarPrimero(), "Deberia borrar el valor booleano true")
}

func TestLista_Vaciada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.BorrarPrimero()

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar el primer elemento")

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "VerPrimero debería causar panico en una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "VerUltimo debería causar panico en una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "BorrarPrimero debería causar panico en una lista vacia")

	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.VerPrimero(), "El primer elemento deberia ser el 2")
	require.EqualValues(t, 2, lista.VerUltimo(), "El ultimo elemento deberia ser el 2")
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía")

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía después de borrar el unico elemento")

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "VerPrimero deberia causar panico en una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "VerUltimo deberia causar panico en una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "BorrarPrimero debería causar panico en una lista vacia")
}

//pruebas del iterador Interno

func TestIterInt_RecorridoCompleto(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(40)
	lista.InsertarUltimo(60)
	lista.InsertarUltimo(80)
	suma := 0
	lista.Iterar(func(num int) bool {
		suma += num
		return true
	})
	require.EqualValues(t, 200, suma, "la suma de los elementos debe ser 200")
}

func TestIterInt_RecorridoCortado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("A")
	lista.InsertarUltimo("B")
	lista.InsertarUltimo("C")
	lista.InsertarUltimo("D")
	lista.InsertarUltimo("E")
	letras := []string{}
	lista.Iterar(func(letra string) bool {
		if letra == "C" {
			return false
		}
		letras = append(letras, letra)
		return true
	})
	require.EqualValues(t, 2, len(letras))
	require.EqualValues(t, "A", letras[0])
	require.EqualValues(t, "B", letras[1])

}

func TestLista_IterarConSuma(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	numeros := []int{1, 2, 3, 4, 5}
	sumaEsperada := 15

	for _, n := range numeros {
		lista.InsertarUltimo(n)
	}

	suma := 0
	lista.Iterar(func(e int) bool {
		suma += e
		return true
	})

	require.Equal(t, sumaEsperada, suma, "la suma de los elementos deberia coincidir con el valor esperado")
}

func TestIterInt_Vacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	elem := 0
	lista.Iterar(func(numero int) bool {
		elem += numero * 2
		return true
	})
	require.EqualValues(t, 0, elem, "no debe haber elementos iterados")
}

func TestIterInt_VaciaConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	llamado := false
	lista.Iterar(func(numero int) bool {
		llamado = true
		return false
	})
	require.False(t, llamado, "no debe llamarse la funcion de iteracion en una lista vacia")
}

func TestIterInt_Buscar(t *testing.T) {
	const letras = "ABCDEFGHIJKLMN"
	incluido := false
	lista := TDALista.CrearListaEnlazada[rune]()
	for _, l := range letras {
		lista.InsertarUltimo(l)
	}
	lista.Iterar(func(letra rune) bool {
		if letra != 'F' {
			return true
		} else {
			incluido = true
			return false
		}
	})
	require.True(t, incluido, "El valor buscado debe estar en la lista")
}

func TestIterInt_EvitarElementos(t *testing.T) {
	cant := 20
	lista := TDALista.CrearListaEnlazada[int]()
	pares := []int{}
	for i := 0; i < cant; i++ {
		lista.InsertarUltimo(i)
	}
	lista.Iterar(func(numero int) bool {
		if numero%2 == 0 {
			pares = append(pares, numero)
		}
		return true
	})
	require.EqualValues(t, cant/2, len(pares), "la mitad de los elementos deberian ser pares")
}

// Pruebas del iterador Externo

func TestIterExt_Recorrer(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 3 {
		lista.InsertarUltimo(i)
	}

	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente(), "Debe haber un elemento al principio")
	require.EqualValues(t, 0, iter.VerActual(), "El iterador debe estar en el primer elemento")

	iter.Siguiente()
	require.True(t, iter.HaySiguiente(), "Debe haber un segundo elemento")
	require.EqualValues(t, 1, iter.VerActual(), "El iterador debe estar en el segundo elemento")

	iter.Siguiente()
	require.True(t, iter.HaySiguiente(), "Debe haber un tercer elemento")
	require.EqualValues(t, 2, iter.VerActual(), "El iterador debe estar en el tercer elemento")

	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), "No debe haber más elementos")

	require.Panics(t, func() {
		iter.VerActual()
	}, "VerActual debería panicar si no hay elemento actual")
}

func TestIterExt_InsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente(), "No debería haber nada para ver")

	iter.Insertar(2.718281828459045)

	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El primer elemento de la lista debe ser el que insertó el iterador")

	require.True(t, iter.HaySiguiente(), "Debería haber algo más para ver")

	iter.Insertar(3.141592653589793)

	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El primer elemento de la lista debe ser el que insertó el iterador")

	iter.Siguiente()
	require.EqualValues(t, 2.718281828459045, iter.VerActual(), "El siguiente elemento debe ser el anterior que insertamos")
}

func TestIterExt_InsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int32]()
	lista.InsertarPrimero(24)
	lista.InsertarUltimo(34)

	iter := lista.Iterador()
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "el elemento actual del iterador debe ser el primero de la lista")
	iter.Siguiente()
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "El elemento actual deberia ser el ultimo")
	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), "No deberia haber algo mas para ver")

	iter.Insertar(47)
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "El elemento actual del iterador debe ser el ultimo de la lista")

	lista.InsertarUltimo(50)
	require.EqualValues(t, 50, lista.VerUltimo(), "El ultimo elemento de la lista debe ser el nuevo insertado")
}

func TestIterExt_InsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(10)
	iter.Siguiente()
	iter.Insertar(20)
	iter.Siguiente()
	iter.Insertar(40)

	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "el elemento actual debe ser 40")

	iter.Insertar(30)
	require.EqualValues(t, 30, iter.VerActual(), "se debe haber insertado el 30")

	iter.Siguiente()
	require.EqualValues(t, 40, iter.VerActual(), "el elemento posterior al 30 debe ser el que era el actual antes de insertar")

	nuevoIter := lista.Iterador()
	require.EqualValues(t, 10, nuevoIter.VerActual(), "el primer elemento debe ser 10")
	nuevoIter.Siguiente()
	require.EqualValues(t, 20, nuevoIter.VerActual(), "el segundo elemento debe ser 20")
	nuevoIter.Siguiente()
	require.EqualValues(t, 30, nuevoIter.VerActual(), "el tercer elemento debe ser 30")
	nuevoIter.Siguiente()
	require.EqualValues(t, 40, nuevoIter.VerActual(), "el cuarto elemento debe ser 40")
}

func TestIterExt_BorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("A")
	lista.InsertarUltimo("B")
	iter := lista.Iterador()

	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El elemento del iterador deberia ser el primero de la lista")

	borrado := iter.Borrar()
	require.EqualValues(t, "A", borrado, "El valor borrado debe ser el que estaba de primero")

	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El primero deberia ser ahora el elemento actual")

	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "El ultimo deberia ser ahora el elemento actual")

	lista.InsertarPrimero("C")
	require.EqualValues(t, "C", lista.VerPrimero(), "El primer elemento debería ser el que se insertó al principio")

	lista.InsertarUltimo("D")
	require.EqualValues(t, "D", lista.VerUltimo(), "El último elemento debería ser el que se insertó al final")
}

func TestIterExt_BorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float32]()
	lista.InsertarUltimo(3.9)
	lista.InsertarPrimero(2.39)
	lista.InsertarPrimero(34.23)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	borrado := iter.Borrar()
	require.EqualValues(t, float32(3.9), borrado, "El elemento borrado deberia ser el ultimo")
	require.EqualValues(t, float32(2.39), lista.VerUltimo(), "El ultimo de la lista deberia ser el anterior al borrado")

	lista.InsertarUltimo(7.77)
	require.EqualValues(t, float32(7.77), lista.VerUltimo(), "El ultimo elemento debe ser el que se inserto al final")
}

func TestIterExt_NoEstaBorrado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)

	iter := lista.Iterador()
	iter.Siguiente()
	borrado := iter.Borrar()
	require.EqualValues(t, 20, borrado, "Debe haberse borrado el 20")

	require.NotEqualValues(t, 20, lista.VerPrimero(), "El primero no debe ser 20")
	require.NotEqualValues(t, 20, lista.VerUltimo(), "El último no debe ser 20")
	require.NotEqualValues(t, 20, iter.VerActual(), "El actual no debe ser 20")

	require.EqualValues(t, 2, lista.Largo(), "La lista debe tener largo 2")

	iterVer := lista.Iterador()
	require.EqualValues(t, 10, iterVer.VerActual())
	iterVer.Siguiente()
	require.EqualValues(t, 30, iterVer.VerActual())

	iterVer.Siguiente()
	require.False(t, iterVer.HaySiguiente(), "No debe haber más elementos")
}

func TestIterExt_VolumenInsertarBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	tam := 10000

	for i := tam - 1; i >= 0; i-- {
		iter.Insertar(i)
		require.EqualValues(t, i, iter.VerActual())
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, tam-1, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
	}

	iter = lista.Iterador()
	for j := 0; j < tam; j++ {
		require.EqualValues(t, j, iter.VerActual())
		require.EqualValues(t, j, lista.VerPrimero())
		require.EqualValues(t, tam-1, lista.VerUltimo())
		require.False(t, lista.EstaVacia())
		require.EqualValues(t, j, iter.Borrar())
	}

	require.True(t, lista.EstaVacia())
}

func TestIterExt_VolumenRecorrido(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 100000
	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
	}

	iter := lista.Iterador()
	for i := 0; i < tam; i++ {
		require.True(t, iter.HaySiguiente(), "HaySiguiente debe ser true en la posición %d", i)
		require.EqualValues(t, i, iter.VerActual(), "El elemento actual del iterador debe ser %d", i)
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente(), "HaySiguiente debe ser false después de recorrer toda la lista")
	require.EqualValues(t, tam, lista.Largo(), "El largo de la lista debe ser %d", tam)
}

func TestIterExt_FinIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), "No debe haber nada mas para ver")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() }, "iter.Siguiente debe causar panico en una lista ya iterada")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() }, "iter.Borrar debe causar panico en una lista ya iterada")
}
