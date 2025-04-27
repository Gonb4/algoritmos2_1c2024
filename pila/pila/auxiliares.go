package pila

/*
func redimensionarSiNecesario[T any](p *pilaDinamica[T]) {
	if p.cantidad == cap(p.datos) {
		nuevoDatos := make([]T, cap(p.datos)*2+1)
		copy(nuevoDatos, p.datos)
		p.datos = nuevoDatos

	} else if p.cantidad*4 <= cap(p.datos) {
		nuevoDatos := make([]T, cap(p.datos)/2)
		copy(nuevoDatos, p.datos)
		p.datos = nuevoDatos
	}
}

func aumentarCapacidad[T any](datos []T) []T {
	nuevoDatos := make([]T, cap(datos)*2)
	copy(nuevoDatos, datos)
	return nuevoDatos
}

func reducirCapacidad[T any](datos []T) []T {
	nuevoDatos := make([]T, cap(datos)/2)
	copy(nuevoDatos, datos)
	return nuevoDatos
}
*/
