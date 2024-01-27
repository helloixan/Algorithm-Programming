package main

import "fmt"

func hitungSkor(soal, skor *int) {
	var waktu int
	*soal, *skor = 0, 0
	for i := 1; i <= 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var (
		peserta, pemenang               string
		jml_soal, nilai, maxsoal, total int
	)

	maxsoal, total = 0, 0
	for selesai := false; !selesai; {
		fmt.Scan(&peserta)
		if peserta == "Selesai" {
			selesai = true
		} else {
			hitungSkor(&jml_soal, &nilai)

			if jml_soal > maxsoal {
				maxsoal = jml_soal
				pemenang = peserta
				total = nilai
			} else if (jml_soal == maxsoal) && (nilai > total) {
				total = nilai
				pemenang = peserta
			}
		}
	}

	fmt.Println(pemenang, maxsoal, total)

}
