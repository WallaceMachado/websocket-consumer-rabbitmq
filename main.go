package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/config"
	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/templates/index.html"))
	t.Execute(w, nil)
}

func operationsInPersonApi(w http.ResponseWriter, r *http.Request) {

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/person-api", operationsInPersonApi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	config.Loader()

	setupRoutes()
}
