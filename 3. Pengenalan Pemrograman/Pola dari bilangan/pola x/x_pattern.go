package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j == i || (j == (n - (i - 1))) {
				fmt.Print(i)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}
