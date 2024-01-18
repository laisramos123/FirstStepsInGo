package main

import (
	"fmt"
	"log"
)

func replace(valueInKelvin int) int {
	return valueInKelvin - 273
}

func main() {
	var celsiusValue, kelvinValue int
	_, err := fmt.Scanln(&kelvinValue)
	if err != nil {
		log.Fatal(err)
	}
	celsiusValue = replace(kelvinValue)

	fmt.Println("O valor em Celcius é:", celsiusValue)
}
