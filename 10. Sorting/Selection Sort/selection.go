package main

import "fmt"

type tab [100]int

func selectionSort(a *tab, n int) {
	var idx_min, temp int

	for i := 0; i < n; i++ {
		idx_min = i
		for j := i + 1; j < n; j++ {
			if a[idx_min] > a[j] {
				idx_min = j
			}
		}
		temp = a[idx_min]
		a[idx_min] = a[i]
		a[i] = temp
	}
}

func binarySearch(t tab, n, x int) bool {
	var kr, kn, med int
	var found bool

	kr = 0
	kn = n - 1
	found = false

	for kr <= kn && !found {
		med = (kr + kn) / 2
		if t[med] < x {
			kr = med + 1
		} else if t[med] > x {
			kn = med - 1
		} else {
			found = true
		}
	}
	return found
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
	var n, x int

	inputData(&t, &n)
	fmt.Println("Data sebelum disort: ")
	cetakArray(t, n)

	fmt.Println("\nData setelah disort: ")
	fmt.Print("Selection sort: ")
	selectionSort(&t, n)
	cetakArray(t, n)

	fmt.Print("\nSearch: ")
	fmt.Scan(&x)
	if binarySearch(t, n, x) {
		fmt.Println("Data ditemukan dengan binary search")
	}

}
