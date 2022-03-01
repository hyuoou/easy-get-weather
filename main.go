package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	owm "github.com/ramsgoli/Golang-OpenWeatherMap"
)

func GetWeatherApi() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	return os.Getenv("WeatherApi")
}

func GetCity() string {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatalln("引数がありません。引数で都市名を選択してください")
	}
	if len(flag.Args()) > 1 {
		log.Fatalln("引数が多すぎます。都市名のみ選択してください")
	}
	return flag.Arg(0)
}

func ToCelsius(temp float64) float64 {
	return (temp - 32) / 1.8
}

func main() {
	api := owm.OpenWeatherMap{API_KEY: GetWeatherApi()}
	cityName := GetCity()

	currentWeather, err := api.CurrentWeatherFromCity(cityName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("都市名  :", currentWeather.Name)
	fmt.Println("天気    :", currentWeather.Weather[0].Main)
	fmt.Printf("気温    : %.2f°C\n", ToCelsius(currentWeather.Temp))
	fmt.Printf("湿度    : %d%%\n", currentWeather.Humidity)
	fmt.Printf("最高気温: %.2f°C\n", ToCelsius(currentWeather.Temp_max))
	fmt.Printf("最低気温: %.2f°C\n", ToCelsius(currentWeather.Temp_min))
}
