package modules

import (
	"encoding/json"
	"net/http"
	"sync"
)

//用於WebAPI:"/api/btc"
//整合不同來源，response json為型別BitCoin的array。
func GetBitCoinUSD(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(3)
	bitcoins := make([]BitCoin, 3, 3)
	go coinMarketCap(bitcoins, &wg)
	go coinGecKo(bitcoins, &wg)
	go nomics(bitcoins, &wg)
	wg.Wait()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	respJSON, _ := json.Marshal(bitcoins)
	w.Write(respJSON)
}
