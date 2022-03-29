/*
	este ejemplo es el mismo que secuencial pero utilizando goroutines como funciones con
	nombre.
	Las goroutinas son hilos propios de go que no utilizan directamente el mecanismo de threading de SO
	son muy livianas en el uso de memoria y en muy eficiente en su ejecución

	para ejecutar una función como goroutine solo se agrega la palabra clave go delante de la función
	ej. 	go CheckOneWebsite(url, workers, &wg)
	tambien se puede realizar con funciones anónimas
	ej.  	go func (u string, workerId int)  {

	Go cuenta con dos grupos de mecanismos de concurrencia.
	el paquete syn que permite forkear goroutines y hacer join de su ejecución
	sync también cuenta con semáforos (mutex y variantes) para proteger variables o "memoria" que
	comparten las goroutines. En este caso las gorutinas no comparten información.
	Por ejemplo, si quisieramos utilizar una slice o map para guardar los resutados que se
	despliegan cada llamada a CheckOneWebsite(...) esa estructura deberia estar protegida
	por un sync.Mutex

	Este ejemplo solamente utiliza el mecanismo de fork y join que se conoce como WaitGroup
	el WaitGroup permite
	  - llevar el contardor de goroutinas a sincronizar Add(n)
	  - decrementar el contador cuando una goroutina termina Done()
	  - esperar que todas las goroutinas en ejecución terminen Wait()

	puede ejecutarlo con
		go run . para ejecutar el main o
		go test -bench . que ejecuta un benchamrk para ver los tiempo que demora la ejecución
		concurrente

	al ejecutar prestar atención al orden en que se despliega el workerId o el url y
	a la cantidad de segundos que duró la ejecución (con go test -bench)


	el ejemplo se adapta de https://github.com/quii/learn-go-with-tests/tree/main/concurrency
	y el libro Concurrency in Go por Katherine Cox-Buday
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

// esta función es la que llama a las goroutinas y las sincroniza
// utilizando sync.WaitGroups
func CheckWebsites(urls []string) {

	var workers int = 0

	// se decalra el grupo de espera
	var wg sync.WaitGroup

	for _, url := range urls {

		wg.Add(1) //se agrega cada goroutina que se ejecuta

		workers++

		go CheckOneWebsite(url, workers, &wg)
	}

	wg.Wait() // se espera a que todas terminen

}

// esta función hace HEAD de el URL e imprime si responde o no
func CheckOneWebsite(url string, workerId int, wg *sync.WaitGroup) {

	defer wg.Done() // cuanto termine defer avisar al WaitGroup

	// el que sigue es el mismo código que en el ejemplo secuencial
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
