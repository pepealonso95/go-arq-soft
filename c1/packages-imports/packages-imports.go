//1

///Packages

/// Todo programa en go esta hecho de paquetes "packages"

/// Todo programa comienza con el package "main"

/// No se puede tener mas de un package por carpeta (con la excepcion del package de testing)
/// Con cambiar las rutas (ponerlos en carpetas distintas) ya son distintos
/// es decir puede haber mas de un package "main"

package main

///Imports

/// Los imports son paqckages importados
/// Puede ser de a varios como aqui o puede ser de a uno
import (
	"fmt"
	"math/rand"

	//Para importar otro paquete en otro directorio
	// hay que usar la direccion del proyecto (declarada en go mod init)
	// y de ahi ir a al directorio del paquete (sin declarar archivos ni paquetes)
	// pongo el nombre del paquete como alias para aclarar codigo, pero no es necesario
	other "go-arq-soft/c1/packages-imports/other-folder"
)

///Nombres exportados
/// "Println" y "Intn" comienzan con mayuscula ya que
/// los metodos con nombres mayuscula son accesibles de afuera del
/// package donde se crearon, de lo contrario no se pueden usar
/// y tendras un codigo "not delcared bu package"

// Todo package debe tener una funcion main para poder correrlo solo
// Solo puede haber una sola funcion main por package
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	// Aqui llamo otra funcion de otro archivo sin importarlo
	HeyFromSomewhereElse()
	// Aqui llamo otra funcion de otra carpeta y otro paquete
	// Lo llamo utilizando el nombre del paquete dentro de la otra carpeta
	other.HeyFromAnotherFolder()
}
