package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/Kontentski/weatherly/config"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server
	})
	ctx           = context.Background()
	cacheDuration = 10 * time.Minute // Cache expiration duration
)

func getWeather(location string) (*Weather, error) {
	cacheKey := fmt.Sprintf("weather:%s", location)
	cachedData, err := rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		//when there is no cached data
		weather, err := fetchWeather(location)
		if err != nil {
			return nil, err
		}

		weatherData, err := json.Marshal(weather)
		if err != nil {
			return nil, err
		}

		rdb.Set(ctx, cacheKey, weatherData, cacheDuration).Err()
		return weather, nil
	} else if err != nil {
		return nil, err
	}

	//when there is cache

	log.Println("Cache hit: data retrieved from cache")

	var weather Weather
	if err := json.Unmarshal([]byte(cachedData), &weather); err != nil {
		return nil, err
	}
	return &weather, nil

}

func fetchWeather(location string) (*Weather, error) {
	config.Init()
	key := config.Config.Key
	url := "https://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + location + "&days=1&aqi=no&alerts=no"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("there is a problem with the url")
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
