# go-api
API is creating weather, currency and crypto info ... 

## Data Sources 

- weather:
 - Igor's flawless weather project (github)[/https://github.com/chubin/wttr.in]
- Currency:
 - Free Api (site)[/www.exchangerate-api.com]
- Crypto:
 - Binance

## Default values
- location = Ankara 
- assets = USD-EUR,USD-GBP,USD-JPY,USD-TRY
- coin = BTCUSDT, ETHUSDT

## Sample queries :

```json
{"time":"2025-01-26T11:04:53Z","weatherbucket":{"status":200,"updatetime":"2025-01-26 10:14 AM","location":"Ankara","temp":"4 °C","weatherDesc":"Partly cloudy","humidity":"87","feelsLikeC":"4","windspeedKm":"4","areaName":"Ankara","latitude":"39.927","longitude":"32.864","country":"Turkey","sunrise":"08:02 AM","sunset":"06:01 PM","moon_illumination":"14","moon_phase":"Waning Crescent","moonrise":"05:48 AM","moonset":"02:30 PM"},"currency":{"status":200,"assets":{"USD-EUR":0.953,"USD-GBP":0.803,"USD-JPY":155.97,"USD-TRY":35.69}},"crypto":{"status":200,"asset":{"BTCUSDT":{"symbol":"BTCUSDT","lastPrice":"104737.15000000","priceChangePercent":"0.131"},"ETHUSDT":{"symbol":"ETHUSDT","lastPrice":"3306.74000000","priceChangePercent":"0.296"}}}}
```
