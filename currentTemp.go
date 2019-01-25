package main

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"fmt"
	"net/http"
	"os"
)

func main() {
	city := "London"
	units := "metric"
	if len(os.Args) > 2 {
		city = os.Args[1]
		units = os.Args[2]
	} else if os.Args[1] == "help" {
			fmt.Printf("Usage examples:\n" +
				"./currentTemp London metric\n" +
				"./currentTemp New York imperial\n" +
				"Defaults are(if no arguments given):\nLondon metric")
			os.Exit(0)
	}
	rootApi := "https://api.openweathermap.org/data/2.5/"
	apiKey := "54c59d8e361e453d8800e52fd9bf54fe"
	resp,err := http.Get(rootApi + "weather?q=" + city + "&appid=" + apiKey + "&units=" + units)
	if err != nil {
		fmt.Printf("Happens")
	}
	defer resp.Body.Close()
	jsonBlob,err := ioutil.ReadAll(resp.Body)
	currentTemp := gjson.Get(string([]byte(jsonBlob)),"main.temp")
	currentWeather := gjson.Get(string([]byte(jsonBlob)),"weather.#.main")
	var currState string
	for _,currentWeather = range currentWeather.Array() {
		currState = currentWeather.String()
	}
	printUnit := "K"
	if units == "metric" {
		printUnit = "C"
	} else if units == "imperial" {
		printUnit = "F"
	}
	fmt.Print(currentTemp.String() + "°" + printUnit + " " + currState)
}
