package modules

import (
	"encoding/json"
	"io/ioutil"
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
		resp, err := http.Get("http://localhost:8080/api/btc")
		if err != nil {
			log.Println("Error create request")
			return
		}
		var bitcoins []BitCoin
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &bitcoins)
		if err := conn.WriteJSON(bitcoins); err != nil {
			log.Println(err)
			return
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
	defer ws.Close()
	log.Println("Client Connected")
	transmitter(ws)
}
