// 2

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Punteros con Metodos:
// Para modificar el valor real de un objeto, el metodo
// debe recibir una referencia al mismo
// Si sacamos el * de Vertex, vean como no afecta el valor
//
// Recuerden que Go infiere el objeto por lo cual
// no es necesario aclarar que queremos acceder a un
// atributo del objeto y no del puntero
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Punteros con Funciones:
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// Noten que go tambien infiere punteros para utilizar metodos
	v.Scale(10)
	fmt.Println(v.X)

	// fmt.Println(v.Abs())
	// Go NO infiere punteros al pasar parametros en funciones
	Scale(&v, 10)
	// fmt.Println(v.Abs())
}
