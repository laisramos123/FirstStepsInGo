package DesafioDio2

import "fmt"

func DesafioDio2Parte2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
		}
		if i%5 == 0 {
			fmt.Println("Pan")
		}
	}
}
