package main

import (
	"context"
	"fmt"
	"math/rand"
	"pubsubexample/pubsub"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	wg := sync.WaitGroup{}
	// Creo el cliente de redis apuntando a nuestra instancia de redis
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	// Chequeo la conexion
	err := client.Ping(context.Background()).Err()
	// Si la conexion falla corto la ejecucion
	if err != nil {
		panic(err)
	}

	// Creo una instancia de publisher
	publisher := pubsub.NewRedisPublisher(client)

	// Creo una instancia de un canal A
	subscriptionChannelA := pubsub.NewRedisSubsciptionChannel(client, "channelA")
	// Creo una instancia de un canal B
	subscriptionChannelB := pubsub.NewRedisSubsciptionChannel(client, "channelB")

	// Subscribo al canal A para que haga un print cuando llegue un mensaje
	cancelA1 := subscriptionChannelA.Subscribe(func(message string) {
		fmt.Println("SubscriberA1:", message)
		// wg.Done()
	})
	// Marco que al final ejecute la funcion para cancelar la subscripcion
	defer cancelA1()

	// Subscribo al canal A para que haga un print distinto cuando llegue un mensaje
	cancelA2 := subscriptionChannelA.Subscribe(func(message string) {
		fmt.Println("SubscriberA2:", message)
	})
	// Marco que al final ejecute la funcion para cancelar la subscripcion
	defer cancelA2()

	// Subscribo al canal B para que haga un print cuando llegue un mensaje
	cancelB := subscriptionChannelB.Subscribe(func(message string) {
		fmt.Println("SubscriberB", message)
		// wg.Done()
	})
	// Marco que al final ejecute la funcion para cancelar la subscripcion
	defer cancelB()

	// Publico un mensaje en el canal A
	wg.Add(1)
	go func() {

		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
			// wg.Add(1)

			publisher.Publish("channelA", fmt.Sprintf("Mensaje A, %d", i))
		}
		wg.Done()
	}()

	// Me aseguro que el wait no se ejecute antes de go func()
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			// wg.Add(1)
			time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second)
			publisher.Publish("channelB", fmt.Sprintf("Mensaje B, %d", i))
		}
		wg.Done()
	}()

	// Espero que termine la ejecucion de todos los waitgroups
	wg.Wait()
	// le doy un poco de tiempo de espera en el que me aseguro que se corra todo
	time.Sleep(500 * time.Millisecond)

	// fmt.Println("Press ENTER to exit")
	// fmt.Scanln()
}
