//2

package main

import "fmt"

func main() {
	// En go una variable puede crearse de forma tradicional
	//explicitando var/const, nombre y tipo
	var i int = 1
	fmt.Println(i)
	// Usualmente se puede inferir el tipo, aqui se infiere
	// que es string ya que se le asigna un string al crearse
	var s = "string"
	fmt.Println(s)
	// Tambien podemos declarar sin poner var usando la notacion
	// de  a := v
	j := 4
	fmt.Println(j)
	// Esto solo genera vars ya que como podemos ver
	// el valor no es constante y se puede modificar
	j = 1
	fmt.Println(j)
}
