package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Get commands
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getCityWeather := getCmd.String("city", "", "Desc: 'get city' will display inputted city")

	// Add commands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCityWeather := addCmd.String("city", "", "Desc: 'add city' will store your city")

	// Saved command
	flag.NewFlagSet("saved", flag.ExitOnError)

	// Checks to see if a subcommand was inputted
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
		fmt.Println("\nYou've inputted an invalid command")
		fmt.Printf("Command \t Description \n")
		fmt.Printf("add     \t displays inputted city \n")
		fmt.Printf("get     \t stores inputted city into local-save \n")
		fmt.Printf("saved   \t displays stored city \n\n")
	}
}

func HandleGet(getCmd *flag.FlagSet, cityweather *string) {
	// Skips the first value(get), reads "-city"
	getCmd.Parse(os.Args[2:])

	if *cityweather == "" {
		fmt.Print("city name is required, use ")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	// Pass string into function, add to API call cityweather
	if *cityweather != "" {
		// Extract Weather object as weatherObject
		weatherObject := getWeather(*cityweather)
		// Determines SIGO and paired note
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
	// Stores string from file in variable
	fileReadCity, fileErr := os.ReadFile("store/storedcity.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	// Check if file is not empty, if it has content then wipe it
	if string(fileReadCity) != "" {
		if err := os.Truncate("store/storedcity.txt", 0); err != nil {
			log.Printf("Failed to truncate: %v\n", err)
		}
	}
	// Calls write function to store new location
	WriteToFile(addCmd, cityweather)
	fmt.Printf("\n%v was stored as your saved city\n\n", *cityweather)
}

func HandleSaved() {
	fileReadCity, fileErr := os.ReadFile("store/storedcity.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	// Stores city name/zip
	zipweather := string(fileReadCity)

	// Check if inputted string is not null
	if zipweather != "" {
		weatherObject := getWeather(zipweather)
		var sigo, note = checkTemp(int(weatherObject.Current.TempF))

		printWeather(weatherObject, sigo, note)
	}
}
