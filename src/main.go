package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Our command name
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// Returns wheather or not you should go outside
	getCityWeather := getCmd.String("city", "", "Desc: Name of your city")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getCityWeather)
	default:
		fmt.Println("You've inputted an invalid command")
	}
}

func HandleGet(getCmd *flag.FlagSet, cityweather *string) {
	getCmd.Parse(os.Args[2:])

	if *cityweather == "" {
		fmt.Print("city is required")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *cityweather != "" {
		locations := getWeather()
		cityweather := *cityweather
		for _, location := range locations {

			/// Check temperature (in degrees fehrenheit)
			var sigo, note = checkTemp(location.Temperature)

			if cityweather == location.City {
				fmt.Printf("\nCity \t Temperature \t SIGO?\n")
				fmt.Printf("%v \t %v \t %v \n", location.City, location.Temperature, sigo)
				fmt.Printf("Note: %v \n\n", note)
			}
		}
	}
}

// Takes in location Object, returns string
func checkTemp(temp int) (sigo string, noteO string) {
	if temp < 10 {
		return "NO", "Stay in your house if you want to survive"
	} else if temp < 20 {
		return "Noooo", "You're going to freeze to death!"
	} else if temp < 35 {
		return "No", "You should put on a scarf"
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
