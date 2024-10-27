package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type thisIsMyStruct struct {
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
	}
	Sys struct {
		Country string `json:"country"`
	}
	Weather struct {
		Main string `json:"main"`
	}
}

var cityName string

func cityInput() string {
	fmt.Print("Enter city name: ")
	fmt.Scan(&cityName)
	return cityName

}

// "https://api.openweathermap.org/data/2.5/weather?q=Rome&appid=631bde274b691a78190999d52bcdd608"
func api() {

	resp, err := http.Get(`https://api.openweathermap.org/data/2.5/weather?q=` + cityInput() + `&appid=631bde274b691a78190999d52bcdd608`)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	callBody, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	data := string(callBody)
	dataByte := []byte(data)
	if err != nil {
		panic(err)

	}

	err = os.WriteFile("output.json", dataByte, 0644)
	if err != nil {
		panic(err)
	}

	readData, err := os.ReadFile("output.json")

	if err != nil {
		panic(err)

	}

	var pointData thisIsMyStruct

	err = json.Unmarshal(readData, &pointData)

	if err != nil {
		return
	}

	fmt.Println(string(readData))
	fmt.Printf("Visibility : %v this is the output data \n", pointData.Visibility)
	fmt.Printf("Wind speed : %v this is the output data 2 \n", pointData.Wind.Speed)
	fmt.Printf("Country: %v this is the output data 3 \n", pointData.Sys.Country)
	fmt.Printf("Weather main: %v this is the output data 4 \n", pointData.Weather.Main)
	// res1 := strings.Split(string(readData), `,"`)
	// fmt.Println(res1)
}
