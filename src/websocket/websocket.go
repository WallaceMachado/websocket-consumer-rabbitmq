package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/queue"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}

func Writer(conn *websocket.Conn) {

	in := make(chan []byte)

	connection := queue.Connect()
	queue.StartConsuming("person_queue", connection, in)

	for payload := range in {
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			fmt.Println(err)
			return
		}
	}

}
