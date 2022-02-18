package main

import (
	"encoding/json"
	"io/ioutil"
)

type location struct {
	Country     string
	State       string
	City        string
	Temperature int
}

// Goes into the JSON file and `gets` the information
func getWeather() (locations []location) {
	fileBytes, err := ioutil.ReadFile("./locations.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &locations)
	if err != nil {
		panic(err)
	}

	return locations
}

// func saveWeather(locations []location) {
// 	locationBytes, err := json.Marshal(locations)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = ioutil.WriteFile("./locations.json", locationBytes, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }
