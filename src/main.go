package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	commands()
}

func commands() {
	// get command flag, gets temperature data based on inputted city
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getCityWeather := getCmd.String("city", "", "Desc: 'get city' will show inputted city")

	// add command flag, saves inputted city into file
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCityWeather := addCmd.String("city", "", "Desc: 'add city' will store your city (Enter zip-code)")

	// saved command, returns saved city
	flag.NewFlagSet("saved", flag.ExitOnError)

	// Checks to see if
	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getCityWeather)
	case "add":
		HandleAdd(addCmd, addCityWeather)
	case "saved":
		HandleSaved()
	default:
		fmt.Println("You've inputted an invalid command")
	}
}

func HandleGet(getCmd *flag.FlagSet, cityweather *string) {
	getCmd.Parse(os.Args[2:])

	if *cityweather == "" {
		fmt.Print("city name is required, use ")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	// Pass string into function, add to API call cityweather
	if *cityweather != "" {
		weatherObject := getWeather(*cityweather)
		var sigo, note = checkTemp(int(weatherObject.Current.TempF))
		printWeather(weatherObject, sigo, note)
	}
}

func HandleAdd(addCmd *flag.FlagSet, cityweather *string) {
	addCmd.Parse(os.Args[2:])

	if *cityweather == "" {
		fmt.Print("city name is required, use ")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
	// Store file values in variable
	fileReadCity, fileErr := os.ReadFile("storedcity.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	// Check if file is not empty, if its not delete contents and continue
	if string(fileReadCity) != "" {
		if err := os.Truncate("storedcity.txt", 0); err != nil {
			log.Printf("Failed to truncate: %v\n", err)
		}
	}
	// Calls pure function, which writes city to file
	WriteToFile(addCmd, cityweather)
	fmt.Printf("\n%v was stored as your saved city\n\n", *cityweather)
}

func HandleSaved() {
	// Read from file
	fileReadCity, fileErr := os.ReadFile("storedcity.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	zipweather := string(fileReadCity)

	// Check if inputted string is not null
	if zipweather != "" {
		weatherObject := getWeather(zipweather)
		var sigo, note = checkTemp(int(weatherObject.Current.TempF))

		printWeather(weatherObject, sigo, note)
	}
}
