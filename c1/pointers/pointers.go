package main

import "fmt"

//Punteros
//Go tiene punteros, un puntero retiene la direccion en memoria de un valor
// El tipo *T es un puntero al valor de tipo T
// El operador & te permite obtener el puntero de un valor

func main() {
	h, i, j := 0, 42, 2701

	a := h         // asigno h a una nueva variable
	a = 1          // la modifico
	fmt.Println(a) // vemos el valor de a
	fmt.Println(h) // y vemos como no afecta h

	p := &i         // puntero de i
	fmt.Println(*p) // lee el valor de i
	*p = 21         // setea un numero con el puntero
	fmt.Println(i)  // lee el nuevo valor de i
	fmt.Println(*p) // lee el valor de i

	p = &j         // puntero de j
	*p = *p / 37   // divide j a traves del puntero
	fmt.Println(j) // lee el nuevo valor de j
}
