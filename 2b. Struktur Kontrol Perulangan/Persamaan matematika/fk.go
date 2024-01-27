package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		k, fk, akar2 float64
	)

	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	fk = math.Pow(((4*k)+2), 2) / (((4 * k) + 1) * ((4 * k) + 3))
	fmt.Println("Nilai f(K) = ", fk)

	/*versi 2*/
	akar2 = (float64(4) / float64(3))
	for i := 1; i <= 94; i++ {
		fk = math.Pow(((4*k)+2), 2) / (((4 * k) + 1) * ((4 * k) + 3))
		akar2 *= fk
	}
	fmt.Println("Nilai akar 2 = ", akar2)
}
