package main

import "fmt"

func main() {
	var (
		num     string
		max_num int
	)

	fmt.Scanln(&num)
	max_num = 0
	for _, single_num := range num {
		if max_num < int(single_num) {
			max_num = int(single_num)
		}
	}
	fmt.Printf("%c", max_num)
}
