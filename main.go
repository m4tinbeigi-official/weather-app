package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Weather structure to hold the API response
type Weather struct {
	Main struct {
		Temp float64 `json:"temp"` // Temperature in Celsius
	} `json:"main"`
	Name string `json:"name"` // City name
}

// Function to fetch weather data from OpenWeatherMap API
func getWeather(city string) (*Weather, error) {
	apiKey := "bd9f0a3dff960116fc8143903981c8ca" // اینجا باید کلید API خود را قرار بدهی
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	// درخواست GET به API ارسال می‌شود
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// خواندن پاسخ API
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// تجزیه داده‌های JSON به ساختار Weather
	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func main() {
	var city string
	fmt.Println("Enter city name:") // از کاربر نام شهر خواسته می‌شود
	fmt.Scanln(&city)

	// فراخوانی تابع برای دریافت وضعیت آب و هوا
	weather, err := getWeather(city)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// نمایش وضعیت آب و هوا
	fmt.Printf("Weather in %s: %.2f°C\n", weather.Name, weather.Main.Temp)
}