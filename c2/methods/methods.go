//1

package main

import (
	"fmt"
	"math"
)

/// En go no hay clases, pero igualmente se pueden
/// definir metodos para tipos de esta forma
type Vertex struct {
	X, Y float64
}

// Es necesario antes del nombre del metodo
// declarar el tipo de la variable a la que aplica
// y como aplica
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Es casi como hacer esto solo que esto se llama
// asi Absoulte(v) en lugar de v.Absolute()
func Absolute(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Tambien podemos declarar metodos para tipos
// que no sean structs solo si los definimos antes
// con un alias dentro del mismo package

type Word string

func (n Word) Arrow() Word {
	return "-> " + n
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	w := Word("hey")

	fmt.Println(w.Arrow())

	//Como no se usaron punteros, el valor de w sigue igual
	fmt.Println(w)

}
