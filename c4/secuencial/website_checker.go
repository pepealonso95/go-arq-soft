/* Ejecución secuencial - este ejemplo es la base para mostrar el uso de concurrencia en los ejemplos de las demás carpetas.
Muestra un slice con urls que se recorre para ver si el llamado a HEAD retorna ok o no.
La función func CheckWebsites(urls []string) recorre el slice y a la función bloqueante CheckOneWebsite(url string,  workerId int)
que despliega en consola un string con true si el sitio responde, false si no lo hace o no existe si el HEAD devuelve error

puede ejecutarlo con:
go run . para ejecutar el main o
go test -bench . que ejecuta un benchamrk para ver los tiempo que demora la ejecución
secuencial

al ejecutar prestar atención al orden en que se despliega el workerId o el url y
a la cantidad de segundos que duró la ejecución (con go test -bench)
*/

package main

import (
	"fmt"
	"net/http"
)

// funcion que chequea si un sitio reponde

func CheckWebsites(urls []string) {
	var workers int = 0
	// recorre el slice de urls y llama a HEAD
	for _, url := range urls {

		workers++
		CheckOneWebsite(url, workers)

	}

}

// esta función hace HEAD de el URL e imprime si responde o no
func CheckOneWebsite(url string, workerId int) {

	response, err := http.Head(url)
	if err != nil {
		fmt.Printf("workerid: %d para url %s no existe \n", workerId, url)

	} else {

		// ok Head sin error, si no es ok retorna url:false si es ok url:true
		if response.StatusCode != http.StatusOK {
			fmt.Printf("workerid: %d para url %s resulta False \n", workerId, url)
		} else {
			fmt.Printf("workerid: %d para url %s resulta True \n", workerId, url)
		}

	}
}

// solo llama a la función de verificar sitios con un slice de urls
func main() {
	var websites = []string{
		"http://ort.edu.uy",
		"http://google.com",
		"http://github.com",
		"http://arqsoft.com",
		"http://netflix.com",
		"http://instagram.com",
		"http://ingsoft.gaston.com",
	}

	fmt.Printf("***** comienzo *****\n")

	CheckWebsites(websites)

	fmt.Printf("***** FIN *****")

}
