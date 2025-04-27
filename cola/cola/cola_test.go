package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaNueva(t *testing.T) { // 1, 5 y 6
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaString := TDACola.CrearColaEnlazada[string]()
	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaFloat64 := TDACola.CrearColaEnlazada[float64]()

	require.True(t, colaInt.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaInt.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaInt.Desencolar() })

	require.True(t, colaString.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaString.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaString.Desencolar() })

	require.True(t, colaBool.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBool.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBool.Desencolar() })

	require.True(t, colaFloat64.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaFloat64.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaFloat64.Desencolar() })
}

func TestEncolar(t *testing.T) { // 2
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaString := TDACola.CrearColaEnlazada[string]()
	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaFloat64 := TDACola.CrearColaEnlazada[float64]()

	require.NotPanics(t, func() { colaInt.Encolar(1) })
	require.False(t, colaInt.EstaVacia())
	require.Equal(t, 1, colaInt.VerPrimero())
	colaInt.Encolar(2)
	require.Equal(t, 1, colaInt.VerPrimero())

	require.NotPanics(t, func() { colaString.Encolar("uno") })
	require.False(t, colaString.EstaVacia())
	require.Equal(t, "uno", colaString.VerPrimero())
	colaString.Encolar("dos")
	require.Equal(t, "uno", colaString.VerPrimero())

	require.NotPanics(t, func() { colaBool.Encolar(true) })
	require.False(t, colaBool.EstaVacia())
	require.Equal(t, true, colaBool.VerPrimero())
	colaBool.Encolar(false)
	require.Equal(t, true, colaBool.VerPrimero())

	require.NotPanics(t, func() { colaFloat64.Encolar(3.14) })
	require.False(t, colaFloat64.EstaVacia())
	require.Equal(t, 3.14, colaFloat64.VerPrimero())
	colaFloat64.Encolar(2.71828)
	require.Equal(t, 3.14, colaFloat64.VerPrimero())
}

func TestDesencolar(t *testing.T) { // 2
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaString := TDACola.CrearColaEnlazada[string]()
	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaFloat64 := TDACola.CrearColaEnlazada[float64]()

	colaInt.Encolar(1)
	colaInt.Encolar(2)
	require.Equal(t, 1, colaInt.Desencolar())
	require.Equal(t, 2, colaInt.VerPrimero())

	colaString.Encolar("hola")
	colaString.Encolar("chau")
	require.Equal(t, "hola", colaString.Desencolar())
	require.Equal(t, "chau", colaString.VerPrimero())

	colaBool.Encolar(false)
	colaBool.Encolar(true)
	require.Equal(t, false, colaBool.Desencolar())
	require.Equal(t, true, colaBool.VerPrimero())

	colaFloat64.Encolar(22.04)
	colaFloat64.Encolar(1.22)
	require.Equal(t, 22.04, colaFloat64.Desencolar())
	require.Equal(t, 1.22, colaFloat64.VerPrimero())
}

func TestDesencolarHastaVacia(t *testing.T) { // 4 y 7
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaString := TDACola.CrearColaEnlazada[string]()
	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaFloat64 := TDACola.CrearColaEnlazada[float64]()

	colaInt.Encolar(444)
	colaInt.Desencolar()
	require.True(t, colaInt.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaInt.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaInt.Desencolar() })

	colaString.Encolar("cuarenta")
	colaString.Desencolar()
	require.True(t, colaString.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaString.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaString.Desencolar() })

	colaBool.Encolar(true)
	colaBool.Desencolar()
	require.True(t, colaBool.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBool.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaBool.Desencolar() })

	colaFloat64.Encolar(0.0)
	colaFloat64.Desencolar()
	require.True(t, colaFloat64.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaFloat64.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaFloat64.Desencolar() })
}

func TestVolumen(t *testing.T) { // 3
	n := 100000
	cola := TDACola.CrearColaEnlazada[int]()

	for i := range n {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}

	for i := 0; i < n-1; i++ {
		cola.Desencolar()
		require.Equal(t, i+1, cola.VerPrimero())
	}

	cola.Desencolar()
	require.True(t, cola.EstaVacia())
}
