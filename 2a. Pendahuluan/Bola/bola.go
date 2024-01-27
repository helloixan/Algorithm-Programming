package main

import (
	"fmt"
	"math"
)

const PI = 3.1415926535

func main() {
	var (
		jejari, volume, luas float64
	)

	fmt.Print("Jejari = ")
	fmt.Scanln(&jejari)

	volume = (float64(4) / float64(3)) * PI * math.Pow(jejari, 3)
	luas = 4 * PI * math.Pow(jejari, 2)

	fmt.Printf("Bola dengan jejari %v memiliki volume %v dan luas kulit %v", jejari, volume, luas)
}
