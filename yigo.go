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
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
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
}

func WriteToFile(addCmd *flag.FlagSet, cityweather *string) {
	err := ioutil.WriteFile("store/storedcity.txt", []byte(*cityweather), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func checkTemp(temp int) (sigo string, noteO string) {
	if temp < 10 {
		return "NO", "Stay in your house if you want to survive"
	} else if temp < 20 {
		return "Noooo", "You're going to freeze to death!"
	} else if temp < 30 {
		return "No", "You should put on a scarf"
	} else if temp < 40 {
		return "Maybe", "It's cold, but if you want to then go for it"
	} else if temp < 50 {
		return "Yes", "Make sure to wear some long clothing"
	} else if temp < 65 {
		return "Yes", "Enjoy the chill weather ;)"
	} else if temp <= 80 {
		return "Yes", "It's nice out, go out there :)"
	} else if temp <= 90 {
		return "Yes", "Beach weather, go get that tan!"
	} else if temp <= 96 {
		return "No", "It's very hot, up to you if you would like to get sweaty"
	} else if temp <= 105 {
		return "NOO", "Our planet will try up soon, do not go outside!"
	}
	return "maybe", "eh"
}
