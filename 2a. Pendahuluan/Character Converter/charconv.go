package main

import "fmt"

func main() {
	var (
		num1, num2, num3, num4, num5 int
		var1, var2, var3             byte
	)

	fmt.Scanln(&num1, &num2, &num3, &num4, &num5)
	fmt.Scanf("%c%c%c", &var1, &var2, &var3)

	var1 += 1
	var2 += 1
	var3 += 1

	fmt.Printf("%c%c%c%c%c", num1, num2, num3, num4, num5)
	fmt.Printf("\n%c%c%c", var1, var2, var3)
}
