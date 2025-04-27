package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func crearNodoCola[T any](elem T) nodoCola[T] {
	return nodoCola[T]{dato: elem}
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := crearNodoCola(elem)

	if c.EstaVacia() {
		c.primero = &nuevoNodo
	} else {
		c.ultimo.prox = &nuevoNodo
	}

	c.ultimo = &nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}

	primerDato := c.primero.dato
	c.primero = c.primero.prox

	if c.EstaVacia() {
		c.ultimo = nil
	}

	return primerDato
}
