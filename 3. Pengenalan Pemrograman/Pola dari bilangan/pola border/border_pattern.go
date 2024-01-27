package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if i == 1 || i == n {
			for j := 1; j <= n; j++ {
				fmt.Print(i, " ")
			}
		} else {
			for j := 1; j <= n; j++ {
				if j == 1 || j == n {
					fmt.Print(i)
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Printf("\n")
	}
}
