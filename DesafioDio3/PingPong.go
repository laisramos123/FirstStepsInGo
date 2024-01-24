package DesafioDio3

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "Ping"
	}
}
func pongar(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "Pong"
	}
}
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func PingPong() {
	var c chan string = make(chan string)

	go pingar(c)
	go pongar(c)
	go imprimir(c)
	var entrada string
	fmt.Scanln(&entrada)
}
