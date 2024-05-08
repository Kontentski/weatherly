package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kontentski/weatherly/config"
)

func getWeather(location string) (*Weather, error) {
	config.Init()
	key := config.Config.Key
	url := "https://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + location + "&days=1&aqi=no&alerts=no"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: status code not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
