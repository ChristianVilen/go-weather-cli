package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	envError := godotenv.Load()
	if envError != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	key := os.Getenv("WEATHER_KEY")
	response, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/helsinki?unitGroup=metric&key=" + key + "&contentType=json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, readError := io.ReadAll(response.Body)
	if readError != nil {
		fmt.Println("Error reading response body")
		fmt.Println(readError.Error())
		os.Exit(1)
	}

	var result Response

	if err := json.Unmarshal(responseData, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	currentTemp := result.CurrentConditions.Temp

	app := &cli.App{
		Name:     "Weather",
		Usage:    "get current weather",
		Version:  "v0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Christian Vilen",
				Email: "c.vilen@outlook.com",
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.Args().Get(0))
			getWeather(currentTemp)
			return nil
		},
	}

	if cliAppErr := app.Run(os.Args); err != nil {
		log.Fatal(cliAppErr)
	}
}
