package main

import "fmt"

func main() {
	var (
		n              int
		daun, max_daun int
	)

	fmt.Scan(&n)
	max_daun = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&daun)
		if daun > max_daun {
			max_daun = daun
		}
	}
	fmt.Println(max_daun)
}
