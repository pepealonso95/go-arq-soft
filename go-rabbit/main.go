package main

import (
	"fmt"
	"sync"
	"time"

	"go-rabbit/workers"
)

func main() {
	wg := sync.WaitGroup{}
	worker, err := workers.BuildRabbitWorker("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	worker.Listen(50, "test-queue", func(message []byte) error {
		fmt.Println(string(message))
		wg.Done()
		return nil
	})

	messageCount := 10
	wg.Add(messageCount)
	for i := 0; i < messageCount; i++ {
		worker.Send("test-queue", []byte("Hello World"))
	}
	wg.Wait()
	worker.Close()

	err = worker.Send("test-queue", []byte("Hello World"))
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
