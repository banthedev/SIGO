package main

import (
	"encoding/json"
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
	fileReadApiKey, fileErr := os.ReadFile("APIKEY.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	api_link := "https://api.weatherapi.com/v1/current.json?key="

	full_link := api_link + string(fileReadApiKey) + "&q=" + cityI + "&aqi=yes"
	response, err := http.Get(full_link)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
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
