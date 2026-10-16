package main

import (
	"log"
	"os"

	"github.com/CodyBense/goeather/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey := os.Getenv("METEOSOURCE_API_KEY")

	weather, _ := api.GetWeather(apiKey)

	api.PrintWeather(weather)
}
