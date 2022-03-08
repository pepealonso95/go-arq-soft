//7

package main

import "fmt"

//Structs
//En go en lugar de clases utilizamos structs

type Vertex struct {
	X int
	Y int
}

func main() {
	// Esta es la notacion del constructor
	v := Vertex{1, 2}
	fmt.Println(v)
	//Accedemos a sus propiedades tal cual objeto
	v.X = 4
	fmt.Println(v.X)
}
