package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		decimal     int
		temp, biner string
	)

	fmt.Scanln(&decimal)

	for decimal > 0 {
		temp += strconv.Itoa(decimal % 2)
		decimal /= 2
	}

	for i := len(temp) - 1; i >= 0; i-- {
		biner += string(temp[i])
	}

	fmt.Println(biner)
	// fmt.Println(temp)
}
