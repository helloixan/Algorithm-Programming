package main

import "fmt"

func factor(n, x int) {
	if x == n {
		fmt.Print(x)
	} else if n%x == 0 {
		fmt.Print(x, " ")
		factor(n, x+1)
	} else {
		factor(n, x+1)
	}
}

func main() {
	var n int

	fmt.Scanln(&n)
	factor(n, 1)
}
