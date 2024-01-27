package main

import "fmt"

func main() {
	var (
		b     int
		prima bool
	)

	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	fmt.Print("Faktor: ")
	prima = true
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			if (i != 1) && (i != b) {
				prima = false
			}
		}
	}
	fmt.Println("Prima: ", prima)
}
