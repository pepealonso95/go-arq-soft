//11

package main

import "fmt"

//Mapas
//En go los mapas son relaciones clave valor
// el valor inicial de un mapa es nil
//la funcion make crea un mapa del tipo que queremos pronto para usar

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
