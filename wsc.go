package core

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/piccuss/wsc/common"
)

var upgrader = websocket.Upgrader{}

//HandlerEcho test for echo sample
func HandlerEcho(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	common.Must(err)
	log.Println("got ws connect: ", conn.RemoteAddr())
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read ws msg err:", err)
			break
		}
		log.Printf("recv ws msg: %s", message)
		//echo test
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write ws msg err:", err)
			break
		}
	}
}
