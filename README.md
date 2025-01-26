# go-api
API is creating weather, currency and crypto info ... 

## Data Sources 

- weather:
 * Igor's flawless weather project (github)[/https://github.com/chubin/wttr.in]
- Currency:
 * Free Api (site)[/www.exchangerate-api.com]
- Crypto:
 * Binance

## Default values
- location = Ankara 
- assets = USD-EUR,USD-GBP,USD-JPY,USD-TRY
- coin = BTCUSDT, ETHUSDT

## Creating docker container :
** Download repo to go-api folder **
```sh
git clone https://github.com/mozkaya1/go-api.git
cd go-api/
```
** Building docker image and running docker container with exposed `:8080` port **

```sh
sudo docker build -t go-api . # Building image 

docker run -it  -p 8080:8080  go-api:latest # Running :8080 port of local machine
Server started on port 8080
```

** Fetching data **
```sh
curl -s  localhost:8080/api|jq  ## you can install jq to better output for json ... 
```

** Response **
```json
{
  "time": "2025-01-26T11:14:29Z",
  "weatherbucket": {
    "status": 200,
    "updatetime": "2025-01-26 01:59 PM",
    "location": "Ankara",
    "temp": "11 °C",
    "weatherDesc": "Sunny",
    "humidity": "58",
    "feelsLikeC": "12",
    "windspeedKm": "4",
    "areaName": "Ankara",
    "latitude": "39.927",
    "longitude": "32.864",
    "country": "Turkey",
    "sunrise": "08:02 AM",
    "sunset": "06:01 PM",
    "moon_illumination": "14",
    "moon_phase": "Waning Crescent",
    "moonrise": "05:48 AM",
    "moonset": "02:30 PM"
  },
  "currency": {
    "status": 200,
    "assets": {
      "USD-EUR": 0.953,
      "USD-GBP": 0.803,
      "USD-JPY": 155.97,
      "USD-TRY": 35.69
    }
  },
  "crypto": {
    "status": 200,
    "asset": {
      "BTCUSDT": {
        "symbol": "BTCUSDT",
        "lastPrice": "104799.99000000",
        "priceChangePercent": "0.287"
      },
      "ETHUSDT": {
        "symbol": "ETHUSDT",
        "lastPrice": "3308.40000000",
        "priceChangePercent": "0.459"
      }
    }
  }
}
```
