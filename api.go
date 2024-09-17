package main

import (
	"bufio"
	"fmt"

	"net/http"
)

func api() {
	fmt.Println("api")

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Rome&appid=631bde274b691a78190999d52bcdd608")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("Response status: ", resp.Status)
	scan := bufio.NewScanner(resp.Body)
	for i := 0; scan.Scan() && i < 5; i++ {
		fmt.Println(scan.Text())

	}
}
