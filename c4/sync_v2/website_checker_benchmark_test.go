// para probar el becnchmark
// go test -bench .
//https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

package main

import (
	"testing"
)

// benchmark de llamadas secuenciales
func BenchmarkWebsiteChecker(b *testing.B) {

	// declara un slice de urls para chequear
	var websites = []string{
		"http://ort.edu.uy",
		"http://google.com",
		"http://github.com",
		"http://arqsoft.com",
		"http://netflix.com",
		"http://instagram.com",
		"http://ingsoft.gaston.com",
	}

	// Demora tenga paciencia!
	// hace el benchmark hasta que N que es un valor que el runtime determina para que sea significativo
	// Demora tenga paciencia!
	//
	for i := 1; i < b.N; i++ {
		CheckWebsites(websites)
	}
}
