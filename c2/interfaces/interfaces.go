// 3
package main

import "fmt"

// Asi se declara una interfaz en go
type I interface {
	M()
}

type T struct {
	S string
}

// Estp quiere decir que T implementa la interfaz I
// no se tiene que declarar explicitamente
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
