/*
	este ejemplo introduce el uso de channels para sincronizar las goroutines.

	Los channels permiten "pasar" datos entre gorutinas y utilizarlos como forma de coordinación

	Este ejemplo se basa en el código de los ejemplos anteriores pero en lugar de imprimir los ruesultados
	del comando http HEAD a consola, se pasan - mediante un canal - nuevamente a la gorutina padre (recordar que el código
	que lanza las gorutinas es también una!)

	ver ejemplos de uso de canales en:
	 - código de ejemplo de varias gorutinas escribiendo de un canal https://go.dev/play/p/quQn7xePLw
	 - código de ejemplo de varias gorutinas leyendo de un canal https://go.dev/play/p/ESq9he_WzS

	al ejecutar prestar atención al orden en que se despliega el workerId o el url y
	a la cantidad de segundos que duró la ejecución (con go test -bench .)

	el ejemplo se adapta de https://github.com/quii/learn-go-with-tests/tree/main/concurrency
	y el libro Concurrency in Go por Katherine Cox-Buday
*/

package main

import (
	"fmt"
	"net/http"
)

// esta función es la que llama a las goroutinas y las sincroniza
// utilizando sync.WaitGroups
func CheckWebsites(urls []string) {

	var workers int = 1

	// se crea un canal que permite pasar strings https://gobyexample.com/channels
	resultStream := make(chan string)

	// recorremos el slice de urls y lanzamos una gorutina para que verifiquen en forma concurrente los
	// sitios. el resultado se pasa por el canal resultStrean.
	// NOTAR que cuando este for termine de ejecutar se sigue ejecutando este código en concurrencia con
	// las gorutinas! o sea, el siguiente for .... se ejecuta concurremente sacando los datos del canal
	for _, url := range urls {

		go CheckOneWebsite(url, workers, resultStream)
		workers++
	}

	//este for se comienza a ejecutar luego de haberse lanzado las gorutinas
	// va sacando del canal a media que las gorutinas agregan
	for i := 0; i < len(urls); i++ {
		// Esta funcion frena el codigo en este lugar hasta que se devuelva un resultado del channel
		// es como que fuerza a que haya un 'await'
		strVal := <-resultStream
		fmt.Printf("%s ", strVal)
	}
	//se cierra el canal porque no se usa mas, todas las gorutinas terminaron y esta
	//ya desplegó todos los datos
	close(resultStream)
}

// esta función hace HEAD de el URL e imprime si responde o no
// notar el último parametro que es un canal de escritura
func CheckOneWebsite(url string, workerId int, results chan<- string) {

	// el que sigue es el mismo código que en el ejemplo secuencial, solo que con un canal
	// las funcion Sprintf crea un string que se pasa al canal
	response, err := http.Head(url)
	if err != nil {
		//Este es el metodo para agregar resultados a un channel,
		// Es importante destacar que los channels son FIFO
		results <- fmt.Sprintf("workerid: %d para url %s no existe \n", workerId, url)
	} else {

		// ok Head sin error, si no es ok retorna url:false si es ok url:true
		if response.StatusCode != http.StatusOK {
			results <- fmt.Sprintf("workerid: %d para url %s resulta False \n", workerId, url)
		} else {
			results <- fmt.Sprintf("workerid: %d para url %s resulta True \n", workerId, url)
		}
	}
}

// solo llama a la función de verificar sitios con un slice de urls
func main() {
	// declara un array de urls para chequear
	var websites = []string{
		"http://ort.edu.uy",
		"http://google.com",
		"http://github.com",
		"http://arqsoft.com",
		"http://netflix.com",
		"http://instagram.com",
		"http://ingsoft.gaston.com",
	}

	fmt.Printf("*****comienzo *****\n")

	CheckWebsites(websites)

	fmt.Printf("***** FIN *****")
}
