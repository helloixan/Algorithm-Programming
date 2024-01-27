package main

import "fmt"

const STOPNUM int = -5313
const NMAX int = 1000

type arr [NMAX]int
type Tabint struct {
	tab arr
	n   int
}

func selectionSort(t *Tabint) {
	var temp, idx_min int

	for i := 0; i < t.n; i++ {
		idx_min = i
		for j := i + 1; j < t.n; j++ {
			if t.tab[idx_min] > t.tab[j] {
				idx_min = j
			}
		}
		temp = t.tab[idx_min]
		t.tab[idx_min] = t.tab[i]
		t.tab[i] = temp
	}
}

func median(t Tabint) int {
	var idx_median int
	if t.n%2 != 0 {
		idx_median = (0 + t.n - 1) / 2
		return t.tab[idx_median]
	} else {
		return (t.tab[(t.n-1)/2] + t.tab[((t.n-1)/2)+1]) / 2
	}
}

func cetak(t Tabint) {
	for i := 0; i < t.n; i++ {
		fmt.Print(t.tab[i], " ")
	}
}

func main() {
	var t Tabint
	var num int

	t.n = 0
	for done := false; !done; {
		fmt.Scan(&num)
		if (num != 0) && (num != STOPNUM) {
			t.tab[t.n] = num
			t.n++
		} else if num == 0 {
			selectionSort(&t)
			// cetak(t)
			fmt.Println(median(t))
		} else {
			done = true
		}
	}
}
