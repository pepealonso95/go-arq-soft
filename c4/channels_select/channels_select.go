// Run codeCopy code
package main

import (
	"fmt"
	"time"
)

func main() {
	//For our example we’ll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)
	//Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "one"
	}()

	// We’ll use select to await both of these values simultaneously,
	// printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			go func() {
				time.Sleep(2 * time.Second)
				c2 <- "two"
			}()
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	close(c1)
	close(c2)
}
