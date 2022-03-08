//9

package main

import "fmt"

//Slices
// Un slice en go es como una lista, es decir
// funciona como un array pero se puede modificar su tamano
// de forma dinamica

//La capacidad se ubtiene utilizando cap(ARRAY)
// El largo se obtiene utilizando len(ARRAY)

func main() {
	// ejemplo de un array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// ejemplo de un slice formado de una parte del array
	var s []int = primes[1:4]
	fmt.Println(s)

	// vemos como agrandamos el slice con otro segmento
	s = primes[2:6]
	fmt.Println(s)

	// Esto es porque no guardan informacion, sino que son como punteros a arrays
	// se puede ver como modifico el slice y afecta el array original
	s[0] = 500
	fmt.Println(s)
	fmt.Println(primes)

	// Se puede extender usando append
	s = append(s, 24, 37)
	fmt.Println(s)

	//Tambien se pueden crear slices utilizando "make"
	m := make([]int, 5)
	fmt.Println(m)

}
