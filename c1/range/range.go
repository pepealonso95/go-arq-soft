//10

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

//Range
// para iterar sobre un slice o map, utilizamos
// la palabre range en un for

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
