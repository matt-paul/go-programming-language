package main

/* this is the shape of our JSON data
{
	"location": "Zzyzx",
	"weather": "sunny",
	"temperature": 30,
	"celsius": true,
	"temp_forecast": [ 27, 25, 28 ],
	"wind": {
		"direction": "NW",
		"speed": 15
	}
}

Important note: Json encoding will only take place on struct fields that are public
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type weatherData struct {
	LocationName string   `json:"dog"`
	Weather      string   `json:"weather"`
	Temperature  int      `json:"temperature"`
	Celsius      bool     `json:"celsius"`
	TempForecast []int    `json:"temp_forecast"`
	Wind         windData `json:"wind"`
}

type windData struct {
	Direction string `json:"direction"`
	Speed     int    `json:"speed"`
}

type loc struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	location := loc{}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}
	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}
	log.Printf("Received: %v\n", location)

	weather := weatherData{
		LocationName: "London",
		Weather:      "Hot hot hot maybe",
		Temperature:  31,
		Celsius:      true,
		TempForecast: []int{30, 32, 29},
		Wind: windData{
			Direction: "S",
			Speed:     20,
		},
	}

	weatherJSON, err := json.Marshal(weather)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(weatherJSON)

}

func server() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":8088", nil)
}

func client() {
	locJSON, err := json.Marshal(loc{Lat: 51.507351, Lon: -0.127758})
	req, err := http.NewRequest("POST", "http://localhost:8088", bytes.NewBuffer(locJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

func main() {
	go server()
	client()
}
