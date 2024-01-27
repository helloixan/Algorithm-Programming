package main

import "fmt"

func star(n int, bintang string) {
	if n == 1 {
		bintang += "*"
		fmt.Println(bintang)
	} else {
		bintang += "*"
		fmt.Println(bintang)
		star(n-1, bintang)
	}
}

func main() {
	var n int

	fmt.Scanln(&n)
	star(n, "")
}
