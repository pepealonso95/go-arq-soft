// 4

package main

import "fmt"

//declaramos func, nombre, parametros y tipo de retorno
func add(x int, y int) int {
	return x + y
}

/// Es posible retornar varios tipos en go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
