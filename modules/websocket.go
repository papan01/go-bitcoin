package modules

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 定義一個發送器，不斷詢問bitcoin api，傳送給websocekt endpoint
func transmitter(conn *websocket.Conn) {
	for {
		req, err := http.NewRequest("GET", "http://localhost:8080/api/btc", nil)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Println("Error create request")
			return
		}
		var bitcoins []BitCoin
		statusCode := sendRequest(req, &bitcoins)
		if statusCode == http.StatusOK {
			if err := conn.WriteJSON(bitcoins); err != nil {
				log.Println(err)
				return
			}
		}
		time.Sleep(time.Second * 20)
	}
}

//websocket:"/ws"
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	go transmitter(ws)
}
