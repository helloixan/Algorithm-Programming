package main

import "fmt"

const NMAX int = 100

type tab [NMAX]int

func insertionSort(a *tab, n int) {
	var temp, j int

	for i := 1; i < n; i++ {
		j = i
		temp = a[j]
		for j > 0 && temp < a[j-1] {
			a[j] = a[j-1]
			j = j - 1
		}
		a[j] = temp
	}
}

func inputData(t *tab, n *int) {
	var data int

	*n = 0
	for selesai := false; !selesai; {
		fmt.Scan(&data)
		if data >= 0 {
			t[*n] = data
			*n++
		} else {
			selesai = true
		}
	}
}

func cetakArray(t tab, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
}

func main() {
	var t tab
	var n int

	inputData(&t, &n)
	fmt.Print("Array sebelum di sort: ")
	cetakArray(t, n)
	insertionSort(&t, n)
	fmt.Print("\nArray sesudah di sort: ")
	cetakArray(t, n)
}
