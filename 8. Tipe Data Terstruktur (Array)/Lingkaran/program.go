package main

import (
	"fmt"
	"math"
)

type lingkaran struct {
	x, y, r float64
}

type titik struct {
	x, y float64
}

func jarak(c lingkaran, p titik) float64 {
	return math.Sqrt(math.Pow(c.x-p.x, 2) + math.Pow(c.y-p.y, 2))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c, p) <= c.r
}

func main() {
	var c1, c2 lingkaran
	var point titik

	fmt.Scanln(&c1.x, &c1.y, &c1.r)
	fmt.Scanln(&c2.x, &c2.y, &c2.r)
	fmt.Scanln(&point.x, &point.y)

	if didalam(c1, point) && didalam(c2, point) {
		fmt.Println("Titik di dalam Lingkaran 1 dan 2")
	} else if didalam(c1, point) {
		fmt.Println("Titik di dalam Lingkaran 1")
	} else if didalam(c2, point) {
		fmt.Println("Titik di dalam Lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
