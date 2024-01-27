package main

import "fmt"

func main() {
	var (
		celcius, reamur, fahrenheit, kelvin float64
	)

	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&celcius)

	reamur = celcius * (float64(4) / float64(5))
	fahrenheit = (celcius * (float64(9) / float64(5))) + 32
	kelvin = (fahrenheit + 459.67) * (float64(5) / float64(9))

	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin: ", int(kelvin))
}
