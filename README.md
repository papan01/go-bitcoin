<h1 align="center">
  Go Bitcoin
</h1>


整合各平台的查詢加密貨幣(Cryptocurrency)API，使用比特幣(BitCoin)作為範例。

## 🌟 Usage

請先確認您有安裝Go在您的環境當中，
首先下載專案:
```
git clone https://github.com/papan01/go-bitcoin.git your-project-name
cd your-project-name
```
在使用前先去[CoinMarketCap](https://coinmarketcap.com/api/documentation/v1/)與[Nomics](https://docs.nomics.com/)申請API使用，並且於`cofing/config_example.json`中填入API Key，接著把`config_example.json`改名為`config.json`。
執行程式:
```
go run main.go
```

打開瀏覽器輸入網址`http://localhost:8080`即可以看到結果。


## ⚔️ Features

- 使用[gorilla/mux](https://www.gorillatoolkit.org/pkg/mux)來完成router。
- 簡單的webscket範例與前端串接，使用[gorilla/websocket](https://www.gorillatoolkit.org/pkg/websocket)。
- 完整的前端程式(使用React、Webpack、eslint等等)。
- limiter限制user訪問與請求的頻率。
- 整合[CoinMarketCap](https://coinmarketcap.com/api/documentation/v1/)、[CoinGecKo](https://www.coingecko.com/api/documentations/v3)與[Nomics](https://docs.nomics.com/)所提供的API(有些需要去申請API Key才能使用)。

## 🕊 API

目前只提供一個簡單的API作為範例使用:
- http://localhost:8080/api/btc
```json
[
  {
    "source_name":"CoinMarketCap",
    "price":9557.25,
    "market_cap":174332454550.96683,
    "volume_24h":51395762662.47019,
    "percent_change_24h":-5.176103215499928
  }
]
```
