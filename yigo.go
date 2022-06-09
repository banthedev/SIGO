package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		WindMph          float64 `json:"wind_mph"`
		Humidity         int     `json:"humidity"`
		FeelsLikeF       float64 `json:"feelslike_f"`
		UV               float64 `json:"uv"`
	} `json:"current"`
}

func getWeather(cityI string) (Weathers Weather) {
	fileReadApiKey, fileErr := os.ReadFile("store/APIKEY.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	api_link := "https://api.weatherapi.com/v1/current.json?key="

	full_link := api_link + string(fileReadApiKey) + "&q=" + cityI + "&aqi=yes"
	response, err := http.Get(full_link)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Weather
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		fmt.Println("error:", err)
	}

	return responseObject
}

func printWeather(Weathers Weather, sigo string, note string) {
	fmt.Printf("\nCITY: %v, %v\nTEMP: %vf / %vc \nSIGO: %v \n", Weathers.Location.Name, Weathers.Location.Country, Weathers.Current.TempF, Weathers.Current.TempC, sigo)
	fmt.Printf("\nNote: %v \n\n", note)
	fmt.Printf("\nFeels like: %v\nUV: %v\nHumidity: %v\nWind Speed(MPH): %v\n", Weathers.Current.FeelsLikeF, Weathers.Current.UV, Weathers.Current.Humidity, Weathers.Current.WindMph)
}

func WriteToFile(addCmd *flag.FlagSet, cityweather *string) {
	err := ioutil.WriteFile("store/storedcity.txt", []byte(*cityweather), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func checkTemp(Weathers Weather) (sigo string, noteO string) {
	// Create score

	// Weather Metrics
	var feels_like_f = Weathers.Current.FeelsLikeF
	var uv = Weathers.Current.UV
	var humidity = Weathers.Current.Humidity
	var wind_speed = Weathers.Current.WindMph
	// score (min = 0, max = 26)
	var score float64 = 0
	// FeelsLikeF: 70-80(S), 50-60 & 80-90(B), 40-50(C), 30-40 & 90-95(D), 20-30(F), 10-20 & 95-100(FF)
	switch {
	case feels_like_f >= 70 && feels_like_f <= 80:
		score += 1
	case feels_like_f >= 50 && feels_like_f < 60:
		score += 2
	case feels_like_f >= 40 && feels_like_f < 50:
		score += 3
	case feels_like_f >= 30 && feels_like_f < 40:
		score += 4
	case feels_like_f >= 90 && feels_like_f < 95:
		score += 4
	case feels_like_f >= 20 && feels_like_f < 30:
		score += 5
	case feels_like_f >= 10 && feels_like_f < 20:
		score += 6
	case feels_like_f >= 0 && feels_like_f <= 10:
		score += 7
	case feels_like_f >= 95 && feels_like_f < 100:
		score += 7
	case feels_like_f < 0 || feels_like_f > 100:
		score += 9
	}
	new_score := score
	fmt.Printf("Feels like Score: %v\n", score)
	// UV: 1-2 (Safe), 3-5(Moderate), 6-7 (High), 8-10(Very High), 11+ (Extreme)
	switch {
	case uv >= 0 && uv <= 2:
		score += 1
	case uv >= 3 && uv <= 5:
		score += 2
	case uv >= 6 && uv <= 7:
		score += 3
	case uv >= 8 && uv <= 10:
		score += 4
	case uv >= 11:
		score += 5
	}
	fmt.Printf("UV Score: %v\n", score-new_score)
	new_score = score
	// Humidity: 30-60 (S), 60-70 & 25-30 (A), >70 & < 25 (F)
	switch {
	case humidity >= 30 && humidity < 60:
		score += 1
	case humidity >= 60 && humidity < 70:
		score += 2
	case humidity >= 25 && humidity < 30:
		score += 2
	case humidity > 70 && humidity < 25:
		score += 3
	}
	fmt.Printf("Humidity Score: %v\n", score-new_score)
	new_score = score
	// Wind speed:
	switch {
	case wind_speed >= 0 && wind_speed < 3: // (Light Air)
		score += 1
	case wind_speed >= 4 && wind_speed < 7: // (Light Breeze)
		score += 2
	case wind_speed >= 8 && wind_speed < 12: // (Gentle Breeze)
		score += 3
	case wind_speed >= 13 && wind_speed < 18: // (Moderate Breeze)
		score += 4
	case wind_speed >= 19 && wind_speed < 24: // (Fresh Breeze)
		score += 5
	case wind_speed >= 25 && wind_speed < 31: // (Strong Breeze)
		score += 6
	case wind_speed >= 32 && wind_speed < 38: // (Moderate Gale)
		score += 7
	case wind_speed >= 39 && wind_speed < 46: // (Fresh Gale)
		score += 8
	case wind_speed >= 47 && wind_speed < 54: // (Strong Gale)
		score += 9
	}
	fmt.Printf("Wind Score: %v\n", score-new_score)
	fmt.Printf("Total Raw score: %v\n", score)
	var scale float64 = 0.0
	scale = (score / 26) * 10

	fmt.Printf("Total Scaled score: %v\n", scale)

	return "Yeee", "Hawwww"

}
