package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char rune

	*n = 0
	for selesai := false; !selesai; {
		fmt.Scanf("%c", &char)
		if (*n <= NMAX) && (char != '.') {
			t[*n] = char
			*n++
		} else {
			selesai = true
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
}

func balikanArray(t *tabel, n int) {
	var temp rune
	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-(i+1)]
		t[n-(i+1)] = temp
	}
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m)
	cetakArray(tab, m)
	balikanArray(&tab, m)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)
}
