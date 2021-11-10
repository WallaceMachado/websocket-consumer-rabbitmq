package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/config"
	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/queue"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conns []*websocket.Conn
var in = make(chan []byte)

func Upgrade(w http.ResponseWriter, r *http.Request) error {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	//gerar um array de ws
	// padrão observer

	conns = append(conns, ws)
	fmt.Println("t ", ws.LocalAddr().String())

	return nil
}

func Writer() {
	// conn vai ser uma range
	// implementar o close

	if len(conns) == 1 {
		connection := queue.Connect()

		queue.StartConsuming(config.RabbitmqQueuePerson, connection, in, "websocket")
	}
	for payload := range in {

		for _, conn := range conns {

			if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
				fmt.Println("err", err)
				return
			}

		}

	}

}
