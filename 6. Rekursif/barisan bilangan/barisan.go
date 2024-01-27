package main

import "fmt"

func barisan(n int) {
	if n == 1 {
		fmt.Print(n, " ")
	} else {
		fmt.Print(n, " ")
		barisan(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var n int

	fmt.Scanln(&n)
	barisan(n)
}
