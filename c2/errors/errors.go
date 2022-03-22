// 4
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// Aqui implementamos la interfaz de 'error'
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// La interfaz 'error' es una interfaz que ya esta definida dentro de go
// Usualmente los errores se devuelven como segundo parametro
// nil en caso que no haya error y sino se devuelve entero
// esto FUERZA el manejo de excepciones y errores
func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return x / x, nil
	} else {
		return 0, &MyError{
			time.Now(),
			"Number cant be negative",
		}
	}
}

func main() {
	// Aca imprimo mis errores
	fmt.Println(Sqrt(2))
	//Aca imprimo mi error custom
	fmt.Println(Sqrt(-2))
}
