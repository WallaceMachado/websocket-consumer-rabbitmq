package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/websocket"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/templates/index.html"))
	t.Execute(w, nil)
}

func OperationsInPersonApi(w http.ResponseWriter, r *http.Request) {

	err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer()
}
