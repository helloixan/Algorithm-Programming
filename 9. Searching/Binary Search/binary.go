package main

import "fmt"

const NMAX int = 100

type arr [NMAX]int
type tabel struct {
	tab arr
	n   int
}

func binarySearch(t tabel, x int) bool {
	var kr, kn, med int
	var found bool

	kr = 0
	kn = t.n - 1
	found = false

	for kr <= kn && !found {
		med = (kr + kn) / 2
		if t.tab[med] < x {
			kr = med + 1
		} else if t.tab[med] > x {
			kn = med - 1
		} else {
			found = true
		}
	}
	return found
}

func sequentialSearch(t tabel, x int) bool {
	var found bool = false
	var idx int = 0

	for !found && idx <= t.n {
		if t.tab[idx] == x {
			found = true
		}
		idx++
	}
	return found
}

func maxValue(t tabel) int {
	var max int = t.tab[0]

	for i := 0; i < t.n; i++ {
		if t.tab[i] > max {
			max = t.tab[i]
		}
	}
	return max
}

func minValue(t tabel) int {
	var min int = t.tab[0]

	for i := 0; i < t.n; i++ {
		if t.tab[i] < min {
			min = t.tab[i]
		}
	}
	return min
}

func inputData(t *tabel) {
	var data int

	t.n = 0
	for selesai := false; !selesai; {
		fmt.Scan(&data)
		if data >= 0 {
			t.tab[t.n] = data
			t.n++
		} else {
			selesai = true
		}
	}
}

func main() {
	var t tabel
	var x int

	inputData(&t)
	fmt.Print("Nilai yang ingin di cari: ")
	fmt.Scan(&x)

	if sequentialSearch(t, x) {
		fmt.Println("Ditemukan dengan sequential search")
	}
	if binarySearch(t, x) {
		fmt.Println("Ditemukan dengan binarysearch")
	}

	fmt.Println("Min Value: ", minValue(t))
	fmt.Println("Max Value: ", maxValue(t))

}
