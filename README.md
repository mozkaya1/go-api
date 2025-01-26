# Weather-Currency-Crypto API
API is creating weather, currency and crypto info ... 

## Data Sources 

- weather:Igor's flawless weather project [@github](https://github.com/chubin/wttr.in)
- Currency: Free Api [@site](https://www.exchangerate-api.com)
- Crypto: Binance

## Default values
- location = Ankara 
- assets = USD-EUR,USD-GBP,USD-JPY,USD-TRY
- coin = BTCUSDT, ETHUSDT

## Creating docker container :
**Download repo to go-api folder**
```sh
git clone https://github.com/mozkaya1/go-api.git
cd go-api/
```
**Building docker image and running docker container with exposed `:8080` port**

```sh
sudo docker build -t go-api . # Building image 

docker run -it  -p 8080:8080  go-api:latest # Running :8080 port of local machine
Server started on port 8080
```

**Fetching data**
```sh
curl -s  localhost:8080/api|jq  ## you can install jq to better output for json ... 
```

**Response**
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
## Sample Queries:
**keys:**
- location="Weather Location"
- assets="Currency Pairs"
- coin="Crypto Pairs"

> All queries support case insensitive
```sh
curl -s  "localhost:8080/api?location=Amsterdam&assets=EuR-TRY,USD-AUD,CHF-USD&coins=ATOMUSDT,SOLUSDT,XRPUSDT"|jq

```
Response
```json
{
  "time": "2025-01-26T11:52:02Z",
  "weatherbucket": {
    "status": 200,
    "updatetime": "2025-01-26 09:55 AM",
    "location": "Amsterdam",
    "temp": "3 °C",
    "weatherDesc": "Partly cloudy",
    "humidity": "87",
    "feelsLikeC": "1",
    "windspeedKm": "9",
    "areaName": "Amsterdam",
    "latitude": "52.374",
    "longitude": "4.890",
    "country": "Netherlands",
    "sunrise": "08:30 AM",
    "sunset": "05:17 PM",
    "moon_illumination": "14",
    "moon_phase": "Waning Crescent",
    "moonrise": "06:59 AM",
    "moonset": "01:12 PM"
  },
  "currency": {
    "status": 200,
    "assets": {
      "CHF-USD": 1.1,
      "EUR-TRY": 37.47,
      "USD-AUD": 1.58
    }
  },
  "crypto": {
    "status": 200,
    "asset": {
      "ATOMUSDT": {
        "symbol": "ATOMUSDT",
        "lastPrice": "6.17100000",
        "priceChangePercent": "-0.660"
      },
      "SOLUSDT": {
        "symbol": "SOLUSDT",
        "lastPrice": "254.64000000",
        "priceChangePercent": "1.665"
      },
      "XRPUSDT": {
        "symbol": "XRPUSDT",
        "lastPrice": "3.12860000",
        "priceChangePercent": "0.488"
      }
    }
  }
}
```


