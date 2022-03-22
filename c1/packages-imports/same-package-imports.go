// dentro de la misma carpeta, si varios archivos comparten
// package, se los trata como el mismo y se pueden llamar entre si
package main

// Los imports no se comparten entre archivos
import (
	"fmt"
)

func HeyFromSomewhereElse() {
	fmt.Printf("Hey")
}
