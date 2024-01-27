package main

import "fmt"

const IDXMAX int = 100

type klub struct {
	nama string
	skor [IDXMAX]int
}

func main() {
	var clubA, clubB klub
	var skorA, skorB, n int

	fmt.Print("Klub A: ")
	fmt.Scan(&clubA.nama)

	fmt.Print("Klub B: ")
	fmt.Scan(&clubB.nama)

	n = 0
	for selesai := false; !selesai; {
		fmt.Printf("Pertandingan %v: ", n+1)
		fmt.Scan(&skorA, &skorB)
		if skorA >= 0 && skorB >= 0 {
			clubA.skor[n] = skorA
			clubB.skor[n] = skorB
			n++
		} else {
			selesai = true
		}
	}

	hitungskor(clubA, clubB, n)
	fmt.Println("Pertandingan selesai")
}

func hitungskor(clubA, clubB klub, n int) {
	for i := 0; i < n; i++ {
		if clubA.skor[i] > clubB.skor[i] {
			fmt.Printf("\nHasil %v: %v", i+1, clubA.nama)
		} else if clubB.skor[i] > clubA.skor[i] {
			fmt.Printf("\nHasil %v: %v", i+1, clubB.nama)
		} else {
			fmt.Printf("\nHasil %v: draw", i+1)
		}
	}
}
