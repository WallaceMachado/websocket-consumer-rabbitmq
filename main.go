package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/config"
	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/router"
)

func main() {
	config.Loader()
	r := router.Generate()

	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r))
}
