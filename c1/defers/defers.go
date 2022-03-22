// 6

package main

import "fmt"

func main() {
	//Defer permite ejecutar el comando cuando la funcion que la engloba complete
	defer fmt.Println("!")

	//Tienen orden LIFO ya que el ultimo en defer se ejecuta primero
	defer fmt.Println("world")

	fmt.Println("hello")
}
