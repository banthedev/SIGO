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
		HandleSavedGet()
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

		/// Check if cityWeather object is the same as the city
		if *cityweather == weatherObject.Location.Name {
			fmt.Printf("\nCity \t Temp \t SIGO? \n")
			fmt.Printf("%v \t %v \t %v \n", weatherObject.Location.Name, weatherObject.Current.TempF, sigo)
			fmt.Printf("\nNote: %v \n\n", note)
		}
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

func HandleSavedGet() {
	// Read from file
	fileReadCity, fileErr := os.ReadFile("storedcity.txt")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	zipweather := string(fileReadCity)
	// (This is temporary), Fetch data
	if zipweather != "" {
		weatherObject := getWeather(zipweather)
		var sigo, note = checkTemp(int(weatherObject.Current.TempF))

		/// Check if cityWeather object is the same as the city
		if zipweather == weatherObject.Location.Name {
			fmt.Printf("\nCity \t Temp \t SIGO? \n")
			fmt.Printf("%v \t %v \t %v \n", weatherObject.Location.Name, weatherObject.Current.TempF, sigo)
			fmt.Printf("\nNote: %v \n\n", note)
		}
	}
}

// Takes in location Object, returns string
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
