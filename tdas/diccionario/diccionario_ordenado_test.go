package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"

	"math/rand"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario ABB vacío no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](CompararInts)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario ABB con un elemento tiene esa clave, únicamente")
	dic := TDADiccionario.CrearABB[string, int](CompararStrings)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](CompararStrings)

	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatos(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)

	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, "miau", dic.Obtener(clave))

	dic.Guardar(clave, "miu")
	require.EqualValues(t, "miu", dic.Obtener(clave))
}

func TestReemplazoDatosMultiples(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean correctos")

	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	for i := 0; i < 500; i++ {
		require.EqualValues(t, 2*i, dic.Obtener(i))
	}
}
func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)

	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
}

func TestReutlizarDeBorrados(t *testing.T) {
	t.Log("Prueba de que no de error al insertar elemento")
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](CompararInts)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestConValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](CompararStrings)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func insertarBalanceado(dic TDADiccionario.Diccionario[string, int], claves []string, valores []int, inicio, fin int) {
	if inicio > fin {
		return
	}
	medio := (inicio + fin) / 2

	dic.Guardar(claves[medio], valores[medio])

	insertarBalanceado(dic, claves, valores, inicio, medio-1)
	insertarBalanceado(dic, claves, valores, medio+1, fin)
}

func pruebaVolumenAbb(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](CompararStrings)

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
	}
	insertarBalanceado(dic, claves, valores, 0, n-1)

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves generadas, " +
		"y que luego podemos borrar sin problemas")
	cant := rand.Intn(5) + 5
	tams := make([]int, cant)
	for i := range cant {
		tams[i] = rand.Intn(500000)
	}
	for _, n := range tams {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				pruebaVolumenAbb(b, n)
			}
		})
	}
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](CompararStrings)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave3 := "Gato"
	clave1 := "Perro"
	clave2 := "Vaca"
	valor3 := "miau"
	valor1 := "guau"
	valor2 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	primero, primer_valor := iter.VerActual()
	require.True(t, dic.Pertenece(primero))
	require.EqualValues(t, primer_valor, dic.Obtener(primero))
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.True(t, primero < segundo)
	require.EqualValues(t, segundo_valor, dic.Obtener(segundo))
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, tercer_valor := iter.VerActual()
	require.True(t, primero < tercero)
	require.True(t, segundo < tercero)
	require.EqualValues(t, tercer_valor, dic.Obtener(tercero))

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](CompararStrings)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter := dic.Iterador()
	iter.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.True(t, primero < segundo)
	require.True(t, primero < tercero)
	require.True(t, segundo < tercero)
	require.True(t, dic.Pertenece(primero))
	require.True(t, dic.Pertenece(segundo))
	require.True(t, dic.Pertenece(tercero))
}

func TestIterInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](CompararStrings)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.True(t, dic.Pertenece(cs[0]))
	require.True(t, dic.Pertenece(cs[1]))
	require.True(t, dic.Pertenece(cs[2]))
	require.True(t, cs[0] < cs[1])
	require.True(t, cs[0] < cs[2])
	require.True(t, cs[1] < cs[2])
}

func TestIterInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](CompararStrings)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestVolumenIterCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](CompararInts)

	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIteradorRangoDesdeHasta(t *testing.T) {
	t.Log("Iterar sobre el diccionario desde un número hasta otro")
	dic := TDADiccionario.CrearABB[int, string](CompararInts)

	dic.Guardar(4, "cuatro")
	dic.Guardar(3, "tres")
	dic.Guardar(2, "dos")
	dic.Guardar(1, "uno")
	dic.Guardar(5, "cinco")

	rangoInferior := 2
	rangoSuperior := 4
	iter := dic.IteradorRango(&rangoInferior, &rangoSuperior)

	require.True(t, iter.HaySiguiente())

	clavesObtenidas := []int{}

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesObtenidas = append(clavesObtenidas, clave)
		iter.Siguiente()
	}

	clavesEsperadas := []int{2, 3, 4}

	require.Equal(t, len(clavesEsperadas), len(clavesObtenidas))

	for i := range clavesEsperadas {
		require.Equal(t, clavesEsperadas[i], clavesObtenidas[i])
	}

	require.False(t, iter.HaySiguiente())
}

func TestIteradorRangoDesde(t *testing.T) {
	t.Log("Iterar sobre el diccionario desde un número en adelante")
	dic := TDADiccionario.CrearABB[int, string](CompararInts)

	dic.Guardar(5, "cinco")
	dic.Guardar(1, "uno")
	dic.Guardar(2, "dos")
	dic.Guardar(4, "cuatro")
	dic.Guardar(3, "tres")

	rangoInferior := 3
	iter := dic.IteradorRango(&rangoInferior, nil)

	require.True(t, iter.HaySiguiente())

	clavesObtenidas := []int{}

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesObtenidas = append(clavesObtenidas, clave)
		iter.Siguiente()
	}

	clavesEsperadas := []int{3, 4, 5}

	require.Equal(t, len(clavesEsperadas), len(clavesObtenidas))

	for i := range clavesEsperadas {
		require.Equal(t, clavesEsperadas[i], clavesObtenidas[i])
	}

	require.False(t, iter.HaySiguiente())
}

func TestIteradorRangoHasta(t *testing.T) {
	t.Log("Iterar sobre el diccionario hasta un número")
	dic := TDADiccionario.CrearABB[int, string](CompararInts)

	dic.Guardar(1, "uno")
	dic.Guardar(2, "dos")
	dic.Guardar(3, "tres")
	dic.Guardar(4, "cuatro")
	dic.Guardar(5, "cinco")

	rangoSuperior := 3
	iter := dic.IteradorRango(nil, &rangoSuperior)

	require.True(t, iter.HaySiguiente())

	clavesObtenidas := []int{}

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesObtenidas = append(clavesObtenidas, clave)
		iter.Siguiente()
	}

	clavesEsperadas := []int{1, 2, 3}

	require.Equal(t, len(clavesEsperadas), len(clavesObtenidas))

	for i := range clavesEsperadas {
		require.Equal(t, clavesEsperadas[i], clavesObtenidas[i])
	}

	require.False(t, iter.HaySiguiente())
}

func TestIteradorInternoRangoDesdeHasta(t *testing.T) {
	t.Log("Iterar sobre el diccionario interno desde un número hasta otro")
	dic := TDADiccionario.CrearABB[int, string](CompararInts)

	dic.Guardar(1, "uno")
	dic.Guardar(2, "dos")
	dic.Guardar(3, "tres")
	dic.Guardar(4, "cuatro")
	dic.Guardar(5, "cinco")

	rangoInferior := 2
	rangoSuperior := 4
	elementosEnRango := []int{2, 3, 4}

	iter := dic.IteradorRango(&rangoInferior, &rangoSuperior)

	clavesObtenidas := []int{}

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesObtenidas = append(clavesObtenidas, clave)
		iter.Siguiente()
	}

	require.Equal(t, len(elementosEnRango), len(clavesObtenidas))

	for i := range elementosEnRango {
		require.Equal(t, elementosEnRango[i], clavesObtenidas[i])
	}

	require.False(t, iter.HaySiguiente())
}

func CompararStrings(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func CompararInts(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
