//1

///Packages

/// Todo programa en go esta hecho de paquetes "packages"

/// Todo programa comienza con el package "main"

package main

///Imports

/// Los imports son paqckages importados
/// Puede ser de a varios como aqui o puede ser de a uno
import (
	"fmt"
	"math/rand"
)

///Nombres exportados
/// "Println" y "Intn" comienzan con mayuscula ya que
/// los metodos con nombres mayuscula son accesibles de afuera del
/// package donde se crearon, de lo contrario no se pueden usar
/// y tendras un codigo "not delcared bu package"

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
