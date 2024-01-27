package main

import "fmt"

func main() {
	var (
		num1    string
		isExist bool = false
		num2    string
	)

	fmt.Scanln(&num1, &num2)
	for _, digit := range num2 {
		if num1 == string(digit) {
			isExist = true
		}
	}
	fmt.Println(isExist)
}
