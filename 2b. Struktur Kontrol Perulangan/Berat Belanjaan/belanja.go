package main

import "fmt"

func main() {
	var (
		total1, total2 float64
		oleng          bool
	)

	// for selesai := false; !selesai; {
	// 	fmt.Print("Masukan berat belanjaan di kedua kantong: ")
	// 	fmt.Scanln(&total1, &total2)
	// 	if (total1 >= 9.0) || (total2 >= 9.0) {
	// 		selesai = true
	// 	}
	// }

	/*Versi modifikasi*/
	for selesai := false; !selesai; {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&total1, &total2)
		if (total1+total2 > 150) || (total1 < 0) || (total2 < 0) {
			selesai = true
		} else {
			oleng = (total1-total2 >= 9) || (total2-total1 >= 9)
			fmt.Println("Sepeda motor pak Andi akan oleng: ", oleng)
		}
	}

	fmt.Println("Proses selesai.")
}
