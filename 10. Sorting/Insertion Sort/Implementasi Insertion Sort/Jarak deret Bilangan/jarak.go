package main

import "fmt"

const NMAX int = 1000

type tab [NMAX]int

func inputData(t *tab, n *int) {
	var data int

	*n = 0
	for done := false; !done; {
		fmt.Scan(&data)
		if data >= 0 {
			t[*n] = data
			*n++
		} else {
			done = true
		}
	}
}

func insertionSort(t *tab, n int) {
	var temp, j int

	for i := 1; i < n; i++ {
		j = i
		temp = t[j]
		for j > 0 && temp < t[j-1] {
			t[j] = t[j-1]
			j--
		}
		t[j] = temp
	}
}

func cekJarak(t tab, n int) int {
	var jarak int

	jarak = t[1] - t[0]
	for i := 2; i < n; i++ {
		if (t[i] - t[i-1]) != jarak {
			jarak = -1
		}
	}
	return jarak
}

func cetakArray(t tab, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
}

func main() {
	var t tab
	var n, jarak int

	inputData(&t, &n)
	insertionSort(&t, n)
	cetakArray(t, n)
	jarak = cekJarak(t, n)
	if jarak > 0 {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}

}
