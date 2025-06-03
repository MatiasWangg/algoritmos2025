package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestCola_alCrearEstaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "Deberia estar vacia al ser creada")
}

func TestCola_verPrimeroPanicsAlEstarVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "VerPrimero deberia causar panico en una cola vacia")
}

func TestCola_desencolarPanicsAlEstarVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Desencolar deberia causar panico en una cola vacia")
}

func TestEncolarDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(0)
	require.False(t, cola.EstaVacia(), "La cola no deberia estar vacia despues de encolar")
	require.Equal(t, 0, cola.VerPrimero(), "El primer valor debería ser el encolado")

	cola.Encolar(1)
	require.Equal(t, 0, cola.VerPrimero(), "El primer valor deberia seguir siendo el mismo")

	desencolado := cola.Desencolar()
	require.Equal(t, 0, desencolado, "Se deberia haber desencolado el primero")
	require.Equal(t, 1, cola.VerPrimero(), "El primer valor ahora deberia ser el segundo encolado")

	desencolado = cola.Desencolar()
	require.Equal(t, 1, desencolado, "Se deberia haber desencolado el ultimo")
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia despues de desencolar todo")
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	tam := 1000

	for i := 0; i < tam; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}
	for j := 0; j < tam; j++ {
		require.Equal(t, j, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia despues de desencolar todo")
}

func TestColaGenerica(t *testing.T) {
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaInt.Encolar(10)
	require.Equal(t, 10, colaInt.Desencolar(), "Deberia desencolar el entero 10")

	colaString := TDACola.CrearColaEnlazada[string]()
	colaString.Encolar("Hola")
	require.Equal(t, "Hola", colaString.Desencolar(), "Deberia desencolar la cadena 'Hola'")

	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaBool.Encolar(true)
	require.Equal(t, true, colaBool.Desencolar(), "Deberia desencolar el valor booleano true")
}
