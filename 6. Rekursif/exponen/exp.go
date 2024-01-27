package main

import "fmt"

func pangkat(x, y int) int {
	if y == 1 {
		return x
	} else {
		return x * pangkat(x, y-1)
	}
}

func main() {
	var x, y int

	fmt.Scanln(&x, &y)
	fmt.Println(pangkat(x, y))
}
