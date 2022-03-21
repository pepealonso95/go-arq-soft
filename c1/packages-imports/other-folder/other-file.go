// dentro de la misma carpeta, si varios archivos comparten
// package, se los trata como el mismo y se pueden llamar entre si
// el package puede llamarse main tambien, no importa
// ya que al estar en otra carpeta se trata como un paquete distinto
package other

// Los imports no se comparten entre archivos
import (
	"fmt"
)

func HeyFromAnotherFolder() {
	fmt.Printf("Hey")
}
