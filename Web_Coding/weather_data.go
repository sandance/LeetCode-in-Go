package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type currentWeather struct {
	Temperature   float64 `json:"temperature"`   // °C
	WindSpeedKmh  float64 `json:"windspeed"`     // km/h
	WindDirection float64 `json:"winddirection"` // degrees
	WeatherCode   int     `json:"weathercode"`
	Time          string  `json:"time"` // ISO-8601
}

type weatherResponse struct {
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	TimeZone       string         `json:"timezone"`
	CurrentWeather currentWeather `json:"current_weather"`
}

func main() {
	//lat := flag.Float64("lat", 37.7749, "latitude")
	//long := flag.Float64("long", -122.4194, "longitude")
	//
	//flag.Parse()
	var lat, long float64

	if _, err := fmt.Scan(&lat, &long); err != nil {
		log.Fatal(err)
	}

	// Build request URL
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		lat, long,
	)

	// call the API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("api returned: %v", resp.Status)
	}

	// Decode json
	var wr weatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&wr); err != nil {
		log.Fatalf("decode failed: %v", err)
	}

	fmt.Println("----- Current Weather -----")
	fmt.Printf("Location      : %.4f, %.4f  (%s)\n", wr.Latitude, wr.Longitude, wr.TimeZone)
	fmt.Printf("Time (ISO)    : %s\n", wr.CurrentWeather.Time)
	//fmt.Printf("Temperature   : %.1f °C  (%.1f °F)\n", c, f)
	fmt.Printf("Wind          : %.1f km/h @ %.0f°\n",
		wr.CurrentWeather.WindSpeedKmh, wr.CurrentWeather.WindDirection)
	fmt.Printf("Weather Code  : %d\n", wr.CurrentWeather.WeatherCode)

}
