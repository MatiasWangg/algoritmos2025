package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPila_alCrearEstaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "Debería estar vacía al ser creada")
}

func TestPila_verTopePanicsAlEstarVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "VerTope debería causar pánico en una pila vacía")
}

func TestPila_desapilarPanicsAlEstarVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Desapilar debería causar pánico en una pila vacía")
}

func TestApilarDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(0)
	require.False(t, pila.EstaVacia(), "La pila no deberia estar vacia")
	require.Equal(t, 0, pila.VerTope(), "El valor del tope debe ser el mismo que el apilado")

	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope(), "El valor del tope debe ser el mismo que el apilado")

	desapilado := pila.Desapilar()
	require.Equal(t, 1, desapilado, "Se tuvo que haber desapilado el ultimo")
	require.Equal(t, 0, pila.VerTope(), "El tope deberia ser 0 despues de haber desapilado")

	desapilado = pila.Desapilar()
	require.Equal(t, 0, desapilado, "Se tuvo que haber desapilado el ultimo")
	require.True(t, pila.EstaVacia(), "La pila esta vacia")

}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	tam := 1000

	for i := 0; i < tam; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for j := tam - 1; j >= 0; j-- {
		require.EqualValues(t, j, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia(), "La pila esta vacia")
}

func TestPilaGenerica(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaInt.Apilar(10)
	require.Equal(t, 10, pilaInt.Desapilar(), "Debería desapilar el entero 10")

	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaString.Apilar("Hola")
	require.Equal(t, "Hola", pilaString.Desapilar(), "Debería desapilar la cadena 'Hola'")

	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaBool.Apilar(true)
	require.Equal(t, true, pilaBool.Desapilar(), "Debería desapilar el valor booleano true")
}
