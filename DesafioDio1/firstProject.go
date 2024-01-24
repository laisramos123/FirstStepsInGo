package main

import (
	"awesomeProject1/DesafioDio2"
	"awesomeProject1/DesafioDio3"
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
	DesafioDio2.DesafioDio2Parte1()
	DesafioDio2.DesafioDio2Parte2()
	DesafioDio3.PingPong()

}
