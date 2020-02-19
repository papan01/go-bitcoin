package modules

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

//CoinMarketCap，https://coinmarketcap.com/api/documentation/v1/
type CoinMarketCap struct {
	Data struct {
		BTC struct {
			Quote struct {
				USD struct {
					Price            float64 `json:"price"`
					MarketCap        float64 `json:"market_cap"`
					Volume24h        float64 `json:"volume_24h"`
					PercentChange24h float64 `json:"percent_change_24h"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"BTC"`
	} `json:"data"`
}

//CoinGecKo，https://www.coingecko.com/api/documentations/v3
type CoinGecKo struct {
	Bitcoin struct {
		Price            float64 `json:"usd"`
		MarketCap        float64 `json:"usd_market_cap"`
		Volume24h        float64 `json:"usd_24h_vol"`
		PercentChange24h float64 `json:"usd_24h_change"`
	} `json:"bitcoin"`
}

//Nomics，https://docs.nomics.com/
type Nomics struct {
	Price     string `json:"price"`
	MarketCap string `json:"market_cap"`
	Interval  struct {
		Volume24h        string `json:"volume"`
		PercentChange24h string `json:"price_change_pct"`
	} `json:"1d"`
}

//最後輸出的結構
type BitCoin struct {
	SourceName       string  `json:"source_name"`
	Price            float64 `json:"price"`
	MarketCap        float64 `json:"market_cap"`
	Volume24h        float64 `json:"volume_24h"`
	PercentChange24h float64 `json:"percent_change_24h"`
}

//儲存最後成功搜尋到的資料
var cache [3]BitCoin

//從CoinMarketCap提供的API，獲取BTC當前價格相關資料，之後轉成我們整合的格式(BitCoin)，append到[]BitCoin中。
func coinMarketCap(t []BitCoin, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Println("Error create request")
		return
	}
	q := url.Values{}
	q.Add("symbol", "BTC")
	q.Add("convert", "USD")
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", viper.GetString("CoinMarketCap-API-KEY"))
	req.URL.RawQuery = q.Encode()
	raw := CoinMarketCap{}
	statusCode := sendRequest(req, &raw)
	if statusCode == http.StatusOK {
		bitcon := BitCoin{
			SourceName:       "CoinMarketCap",
			Price:            raw.Data.BTC.Quote.USD.Price,
			MarketCap:        raw.Data.BTC.Quote.USD.MarketCap,
			Volume24h:        raw.Data.BTC.Quote.USD.Volume24h,
			PercentChange24h: raw.Data.BTC.Quote.USD.PercentChange24h,
		}
		t[0] = bitcon
		cache[0] = bitcon
	} else if cache[0] != (BitCoin{}) {
		t[0] = cache[0]
	}
}

//從CoinGecKo提供的API，獲取BTC當前價格相關資料，之後轉成我們整合的格式(BitCoin)，append到[]BitCoin中。
func coinGecKo(t []BitCoin, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/simple/price", nil)
	if err != nil {
		log.Println("Error create request")
		return
	}
	q := url.Values{}
	q.Add("ids", "bitcoin")
	q.Add("vs_currencies", "usd")
	q.Add("include_market_cap", "true")
	q.Add("include_24hr_vol", "true")
	q.Add("include_24hr_change", "true")
	req.URL.RawQuery = q.Encode()
	raw := CoinGecKo{}
	statusCode := sendRequest(req, &raw)
	if statusCode == http.StatusOK {
		bitcon := BitCoin{
			SourceName:       "CoinGecKo",
			Price:            raw.Bitcoin.Price,
			MarketCap:        raw.Bitcoin.MarketCap,
			Volume24h:        raw.Bitcoin.Volume24h,
			PercentChange24h: raw.Bitcoin.PercentChange24h,
		}
		t[1] = bitcon
		cache[1] = bitcon
	} else if cache[1] != (BitCoin{}) {
		t[1] = cache[1]
	}
}

//從Nomics提供的API，獲取BTC當前價格相關資料，之後轉成我們整合的格式(BitCoin)，append到[]BitCoin中。
func nomics(t []BitCoin, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest("GET", "https://api.nomics.com/v1/currencies/ticker?key="+viper.GetString("Nomics-API-KEY")+"&ids=BTC&interval=1d", nil)
	if err != nil {
		log.Println("Error create request")
		return
	}
	var raw []Nomics
	statusCode := sendRequest(req, &raw)
	if statusCode == http.StatusOK {
		price, _ := strconv.ParseFloat(raw[0].Price, 64)
		marketCap, _ := strconv.ParseFloat(raw[0].MarketCap, 64)
		volume24h, _ := strconv.ParseFloat(raw[0].Interval.Volume24h, 64)
		percentChange24h, _ := strconv.ParseFloat(raw[0].Interval.PercentChange24h, 64)
		bitcon := BitCoin{
			SourceName:       "Nomics",
			Price:            price,
			MarketCap:        marketCap,
			Volume24h:        volume24h,
			PercentChange24h: percentChange24h,
		}
		t[2] = bitcon
		cache[2] = bitcon
	} else if cache[2] != (BitCoin{}) {
		t[2] = cache[2]
	}
}

// 發出請求，給據給定的http.Request，回傳http status code，並且將得到的結果unmarshal到raw中
func sendRequest(r *http.Request, raw interface{}) int {
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("Error sending request to server")
		return resp.StatusCode
	}
	reqBody, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(reqBody, raw); err != nil {
		panic(err)
	}
	return resp.StatusCode
}
