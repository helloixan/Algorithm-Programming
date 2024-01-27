package main

import "fmt"

func main() {
	var (
		biayakg, biayagr, total_biaya float64
		berat, beratkg, sisaberat     int
	)

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	beratkg = berat / 1000
	sisaberat = berat % 1000
	biayakg = float64(beratkg) * 10000

	if (beratkg > 10) && (sisaberat < 1000) {
		biayagr = 0
	} else if sisaberat >= 500 {
		biayagr = float64(5 * sisaberat)
	} else {
		biayagr = float64(15 * sisaberat)
	}

	total_biaya = biayakg + biayagr

	fmt.Printf("Detail berat: %v kg + %v gr", beratkg, sisaberat)
	fmt.Printf("\nDetail Biaya: Rp. %v + Rp. %v", biayakg, biayagr)
	fmt.Printf("\nTotal biaya: Rp. %v", total_biaya)
}
