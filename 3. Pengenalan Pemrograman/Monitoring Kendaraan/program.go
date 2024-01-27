package main

import "fmt"

func main() {
	var (
		kendaraan           byte
		tipeA, tipeB, tipeC int = 0, 0, 0
	)

	for selesai := false; !selesai; {
		fmt.Scan(&kendaraan)
		if kendaraan == 'A' {
			tipeA += 1
		} else if kendaraan == 'B' {
			tipeB += 1
		} else if kendaraan == 'C' {
			tipeC += 1
		} else {
			selesai = true
		}
	}

	fmt.Printf("\nTipe A = %c", tipeA)
	fmt.Printf("\nTipe B = %c", tipeB)
	fmt.Printf("\nTipe C = %c", tipeC)

}
