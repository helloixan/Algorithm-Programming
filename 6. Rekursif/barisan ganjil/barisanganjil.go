package main

import "fmt"

func barisan(n, i int) {
	if i <= n {
		fmt.Print(i, " ")
		barisan(n, i+2)
	}
}

func main() {
	var n int

	fmt.Scanln(&n)
	barisan(n, 1)
}
