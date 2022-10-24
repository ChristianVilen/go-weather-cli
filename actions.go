package main

import (
	"fmt"
	"math"
	"strings"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func formatFloat(num float64) string {
	s := fmt.Sprintf("%.4f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func getWeather(currentTemp float64) {
	s := "current temperature in helsinki: " + formatFloat(toFixed(currentTemp, 0))
	fmt.Println(s)
}
