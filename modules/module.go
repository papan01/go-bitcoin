package modules

import "net/http"

type Module struct {
	APIMap map[string]func(http.ResponseWriter, *http.Request)
}

func (t *Module) Initialize() {
	t.APIMap = make(map[string]func(http.ResponseWriter, *http.Request))
	//BitCoinApi
	t.APIMap["/btc"] = GetBitCoinUSD
}
