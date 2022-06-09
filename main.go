package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

/**
	Title:	Main Function

	Description: This is the main driving function

	Features:
	- Creates flags(get, add,saved) and its subcommands(-city)
	- Reads input and checks if proper command was inputed
	- Uses switch statement to call selected command function
**/
func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getCityWeather := getCmd.String("city", "", "Desc: 'get city' will display inputted city")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCityWeather := addCmd.String("city", "", "Desc: 'add city' will store your city")

	flag.NewFlagSet("saved", flag.ExitOnError)

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

/**
	Title:	HandleGet

	Description: Function for "get" command, retrieves weather object given inputted city(string)

	Features:
	- Parses definitions that should not be in the argument list
	- Checks if city inputted is valid
	- If inputted city is not empty and a string, call getWeather and return information
		- Create two variables which are sent into function call checkTemp
		  these variables are then sent to be printed by pure function
**/
func HandleGet(getCmd *flag.FlagSet, cityweather *string) {
	getCmd.Parse(os.Args[2:])

	if *cityweather == "" {
		fmt.Print("city name is required, use ")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *cityweather != "" {
		weatherObject := getWeather(*cityweather)
		var sigo, note = checkTemp(weatherObject)
		printWeather(weatherObject, sigo, note)
	}
}

/**
	Title:	HandleAdd

	Description: Function for "add" command, stores inputted city(string) to txt file

	Features:
	- Parses definitions that should not be in the argument list
	- Checks if city inputted is valid
	- Reads "storedcity.txt" and checks if its not empty. If it has a city
	  it will use Truncate() to remove the string. (This is prevent multiple
	  cities being in the file)
	- Calls WriteToFile() which stores the string in the file
	- Prints that storing is complete
**/
func HandleAdd(addCmd *flag.FlagSet, cityweather *string) {
	addCmd.Parse(os.Args[2:])

	if *cityweather == "" {
		fmt.Print("city name is required, use ")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
	fileReadCity, fileErr := os.ReadFile("store/storedcity.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	if string(fileReadCity) != "" {
		if err := os.Truncate("store/storedcity.txt", 0); err != nil {
			log.Printf("Failed to truncate: %v\n", err)
		}
	}
	WriteToFile(addCmd, cityweather)
	fmt.Printf("\n%v was stored as your saved city\n\n", *cityweather)
}

/**
	Title:	HandleSaved

	Description: Function for "saved" command, returns saved city information to user

	Features:
	- Reads "storedcity.txt"
	- Create string object of text(the city) in the file
	- Checks if string is not empty, and calls getWeather() and checkTemp(0)
**/
func HandleSaved() {
	fileReadCity, fileErr := os.ReadFile("store/storedcity.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	zipweather := string(fileReadCity)

	if zipweather != "" {
		weatherObject := getWeather(zipweather)
		// var sigo, note = checkTemp(int(weatherObject.Current.TempF))
		var sigo, note = checkTemp(weatherObject)
		printWeather(weatherObject, sigo, note)
	}
}
