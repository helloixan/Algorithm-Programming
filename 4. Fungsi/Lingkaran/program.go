package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cx2, cy1, cy2, x, y, r1, r2 float64

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	if didalam(cx1, cy1, r1, x, y) && didalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(cx1, cy1, r1, x, y) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func jarak(a, b, c, d float64) float64 {
	return (math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2)))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
