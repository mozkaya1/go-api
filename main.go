package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type All struct {
	Time          string        `json:"time"`
	WeatherBucket WeatherBucket `json:"weatherbucket"`
	Currency      Currency      `json:"currency"`
	Crypto        Crypto        `json:"crypto"`
}

type WeatherBucket struct {
	Status            int    `json:"status"`
	UpdateTime        string `json:"updatetime"`
	Location          string `json:"location"`
	Temperature       string `json:"temp"`
	WeatherDesc       string `json:"weatherDesc"`
	Humidity          string `json:"humidity"`
	FeelsLikeC        string `json:"feelsLikeC"`
	WindspeedKmph     string `json:"windspeedKm"`
	NearestArea       string `json:"areaName"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Country           string `json:"country"`
	Sunrise           string `json:"sunrise"`
	Sunset            string `json:"sunset"`
	Moon_illumination string `json:"moon_illumination"`
	Moon_phase        string `json:"moon_phase"`
	Moonrise          string `json:"moonrise"`
	Moonset           string `json:"moonset"`
}

type Currency struct {
	Status int                `json:"status"`
	Assets map[string]float64 `json:"assets"`
}

type Crypto struct {
	Status int                    `json:"status"`
	Assets map[string]CryptoAsset `json:"asset"`
}

type CryptoAsset struct {
	Symbol             string `json:"symbol"`
	LastPrice          string `json:"lastPrice"`
	PriceChangePercent string `json:"priceChangePercent"`
}

func GetCrypto(coins []string) (Crypto, error) {
	query := fmt.Sprintf(`["%s"]`, strings.Join(coins, `","`))
	encoded_query := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/24hr?symbols=%s", encoded_query)
	resp, err := http.Get(url)
	if err != nil {
		return Crypto{}, errors.New("Failed to Fetch Crypto")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Crypto{}, errors.New("Failed to read Crypto Response Body")
	}

	var Assets []CryptoAsset

	err = json.Unmarshal(body, &Assets)
	if err != nil {
		return Crypto{}, errors.New("failed to parse weather JSON")
	}

	NewData := Crypto{
		Status: 200,
		Assets: make(map[string]CryptoAsset),
	}
	for _, k := range Assets {
		NewData.Assets[k.Symbol] = k
	}

	return NewData, nil
}

// Fetch weather de"tails
func GetWeather(location string) (WeatherBucket, error) {
	url := fmt.Sprintf("http://wttr.in/%s?format=j1", location)
	resp, err := http.Get(url) // Use JSON format from wttr.in
	if err != nil {
		return WeatherBucket{}, errors.New("failed to fetch weather")
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherBucket{}, errors.New("failed to read weather response body")
	}

	// Parse the JSON into WeatherBucket
	var weatherData struct {
		CurrentCondition []struct {
			WeatherDesc []struct {
				Value string `json:"value"`
			} `json:"weatherDesc"`
			TempC      string `json:"temp_C"`
			Windspeed  string `json:"windspeedKmph"`
			Humidity   string `json:"humidity"`
			FeelsLikeC string `json:"FeelsLikeC"`
			Updatetime string `json:"localObsDateTime"`
		} `json:"current_condition"`
		Area []struct {
			AreaName []struct {
				Value string `json:"value"`
			} `json:"areaName"`
			Country []struct {
				Value string `json:"value"`
			} `json:"country"`
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"nearest_area"`
		Weather []struct {
			Astronomy []struct {
				Sunrise           string `json:"sunrise"`
				Sunset            string `json:"sunset"`
				Moon_illumination string `json:"moon_illumination"`
				Moon_phase        string `json:"moon_phase"`
				Moonrise          string `json:"moonrise"`
				Moonset           string `json:"moonset"`
			} `json:"astronomy"`
		} `json:"weather"`
	}
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return WeatherBucket{}, errors.New("failed to parse weather JSON")
	}

	// Construct the WeatherBucket
	weather := WeatherBucket{
		Status:            200,
		Location:          location,
		Temperature:       weatherData.CurrentCondition[0].TempC + " °C",
		WeatherDesc:       weatherData.CurrentCondition[0].WeatherDesc[0].Value,
		Humidity:          weatherData.CurrentCondition[0].Humidity,
		FeelsLikeC:        weatherData.CurrentCondition[0].FeelsLikeC,
		UpdateTime:        weatherData.CurrentCondition[0].Updatetime,
		WindspeedKmph:     weatherData.CurrentCondition[0].Windspeed,
		NearestArea:       weatherData.Area[0].AreaName[0].Value,
		Latitude:          weatherData.Area[0].Latitude,
		Longitude:         weatherData.Area[0].Longitude,
		Country:           weatherData.Area[0].Country[0].Value,
		Sunrise:           weatherData.Weather[0].Astronomy[0].Sunrise,
		Sunset:            weatherData.Weather[0].Astronomy[0].Sunset,
		Moon_illumination: weatherData.Weather[0].Astronomy[0].Moon_illumination,
		Moon_phase:        weatherData.Weather[0].Astronomy[0].Moon_phase,
		Moonrise:          weatherData.Weather[0].Astronomy[0].Moonrise,
		Moonset:           weatherData.Weather[0].Astronomy[0].Moonset,
	}

	return weather, nil
}

// Uniq list for base assets
func Uniq(b []string, value string) []string {
	for _, k := range b {
		if k == value {
			return b
		}
	}
	return append(b, value)
}

// Fetch currency details
func GetCurrency(filterAssets []string) (Currency, error) {
	// mq = strings.ToUpper(mq)
	// if mq == "" {
	// 	mq = "USD"

	NewData := make(map[string]float64)
	filteredAssets := make(map[string]float64)
	Base := make([]string, 0)
	if len(filterAssets) > 0 {
		for _, k := range filterAssets {
			line := strings.Split(k, "-")
			Base = Uniq(Base, line[0])
		}
	} else {
		Base = []string{"USD"}
	}

	for _, k := range Base {
		url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s", k)
		resp, err := http.Get(url) // Example API
		if err != nil {
			return Currency{}, errors.New("failed to fetch currency data")
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return Currency{}, errors.New("failed to read currency response body")
		}

		// Parse the JSON into Currency
		var currencyData struct {
			Rates map[string]float64 `json:"rates"`
		}
		err = json.Unmarshal(body, &currencyData)
		if err != nil {
			return Currency{}, errors.New("failed to parse currency JSON")
		}

		for key, value := range currencyData.Rates {
			new_key := fmt.Sprintf("%s-%s", k, key)
			NewData[new_key] = value
		}
	}
	// Construct the Currency struct // exist == True then assign filtered assets to return.

	if len(filterAssets) > 0 {
		for _, asset := range filterAssets {
			if rate, exists := NewData[asset]; exists {
				filteredAssets[asset] = rate
			}
		}
	} else {
		// No filtering, return all assets
		filteredAssets = map[string]float64{
			"USD-TRY": NewData["USD-TRY"],
			"USD-EUR": NewData["USD-EUR"],
			"USD-GBP": NewData["USD-GBP"],
			"USD-JPY": NewData["USD-JPY"],
		}
	}

	// Construct the Currency struct
	currency := Currency{
		Status: 200,
		Assets: filteredAssets,
	}
	return currency, nil
}

// HTTP handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Set CORS headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length,Content-Range")
	w.Header().Set("Vary", "Origin")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Length", 0)
	w.Header().Set("Content-Type", "application/json")

	// Get the current weather and currencies

	location := r.URL.Query().Get("location")
	if location == "" {
		location = "Ankara"
	}

	assets := r.URL.Query().Get("assets")
	var filterAssets []string
	if assets != "" {
		filterAssets = strings.Split(strings.ToUpper(assets), ",") // Split comma-separated assets
	}

	weather, err := GetWeather(location)
	if err != nil {
		// http.Error(w, "Could not fetch weather data", http.StatusInternalServerError)
		http.Error(w, "", http.StatusInternalServerError)
		weather = WeatherBucket{
			Status: 500,
		}
		// return
	}

	currency, err := GetCurrency(filterAssets)
	if err != nil {
		http.Error(w, "Could not fetch currency data", http.StatusInternalServerError)
		return
	}

	coins := r.URL.Query().Get("coins")
	var coin []string
	if coins == "" {
		coin = []string{"BTCUSDT", "ETHUSDT"}
	} else {
		coin = strings.Split(strings.ToUpper(coins), ",")
	}

	crypto, err := GetCrypto(coin)
	// Construct the response
	response := All{
		Time:          time.Now().Format(time.RFC3339),
		WeatherBucket: weather,
		Currency:      currency,
		Crypto:        crypto,
	}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api", helloHandler)
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
