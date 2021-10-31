package main

import (
	"fmt"

	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/config"
	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/queue"
)

func main() {

	config.Loader()

	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming("person_queue", connection, in)

	for payload := range in {
		fmt.Println(string(payload))
	}

}
