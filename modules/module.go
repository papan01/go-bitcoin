package modules

import "net/http"

type Module struct {
	RoutesMap map[string]func(http.ResponseWriter, *http.Request)
}

func (t *Module) Initialize() {
	t.RoutesMap = make(map[string]func(http.ResponseWriter, *http.Request))
	//pages
	t.RoutesMap["/"] = Index
	//bitCoin
	t.RoutesMap["/api/btc"] = GetBitCoinUSD
}
