package server

import (
	"bitcoin/modules"

	"github.com/gorilla/mux"
)

func registerHandlers(router *mux.Router) {
	//pages
	router.HandleFunc("/", modules.Index)
	//api
	s := router.PathPrefix("/api").Subrouter()
	s.HandleFunc("/btc", modules.GetBitCoinUSD).Methods("GET")
	//websocket
	router.HandleFunc("/ws", modules.WsEndpoint)
}
