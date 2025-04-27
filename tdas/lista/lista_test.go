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

func TestLista_Volumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 1000

	for i := 0; i < tam; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, 0, lista.VerPrimero())
	}
	for j := 0; j < tam; j++ {
		require.Equal(t, j, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todo")
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

func TestIterInt_Volumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 10000
	for i := range tam {
		lista.InsertarUltimo(i)
	}
	lista.Iterar(func(e int) bool {
		return true
	})
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
	lista := TDALista.CrearListaEnlazada[int]()
	pares := []int{}
	for i := range 20 {
		lista.InsertarUltimo(i)
	}
	lista.Iterar(func(numero int) bool {
		if numero%2 == 0 {
			pares = append(pares, numero)
		}
		return true
	})
	require.EqualValues(t, len(pares), lista.Largo()/2, "La mitad de elementos de la lista son pares")
}

// Pruebas del iterador Externo

func TestIterExt_Recorrer(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 3 {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	require.EqualValues(t, 0, iter.VerActual(), "el iterador debe estar en el primer elemento")
	iter.Siguiente()
	require.EqualValues(t, 1, iter.VerActual(), "el iterador debe estar en el segundo elemento")
	iter.Siguiente()
	require.EqualValues(t, 2, iter.VerActual(), "el iterador debe estar en el tercer elemento")
	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), "No debe haber nada mas para ver")
}
func TestIterExt_InsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente(), "No deberia haber nada para ver")
	iter.Insertar(2.718281828459045235360)
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual(), "El valor primer elemento de la lista debe ser el que inserto el iterador")
	require.True(t, iter.HaySiguiente(), "Deberia haber algo mas para ver")
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
	require.False(t, iter.HaySiguiente(), " No deberia haber algo mas para ver")
	iter.Insertar(47)
	require.EqualValues(t, lista.VerUltimo(), iter.VerActual(), "El elemento actual del iterador debe ser el ultimo de la lista")
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

}

func TestIterExt_NoEstaBorrado(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Borrar()
	require.NotEqualValues(t, 20, lista.VerPrimero())
	require.NotEqualValues(t, 20, lista.VerUltimo())
	require.NotEqualValues(t, 20, iter.VerActual())
	require.EqualValues(t, 2, lista.Largo())

}

func TestIterExt_VolumenInsertarBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	tam := 10000
	for i := tam; i >= 0; i-- {
		iter.Insertar(i)
		require.EqualValues(t, i, iter.VerActual(), "El valor actual de la iteracion deberia ser j")
	}
	for j := range tam {
		require.EqualValues(t, j, iter.VerActual(), "El valor actual de la iteracion deberia ser j")
		require.EqualValues(t, j, iter.Borrar(), "El valor borrado debe ser j")
	}
}

func TestIterExt_VolumenRecorrido(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	tam := 100000
	for i := range tam {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	j := 0
	for iter.HaySiguiente() {
		require.EqualValues(t, j, iter.VerActual(), "el elemento actual del iterador debe ser el mismo que el contador j")
		iter.Siguiente()
		j++
	}
	require.EqualValues(t, j, lista.Largo(), "el contador j debe ser igual al largo total de la lista")
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