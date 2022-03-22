// 5

package main

import "fmt"

func main() {
	sum := 1
	// un for en go se escribe con ; y sin parentesis
	// la declaracion de la variable y el post son opcionales
	for i := 0; i < 10; i++ {
		sum += sum
	}
	fmt.Println(sum)
	//en go el for es while tambien
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
