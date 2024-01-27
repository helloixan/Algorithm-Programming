package main

import "fmt"

func main() {
	var (
		success                bool
		tab1, tab2, tab3, tab4 string
	)

	success = true
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %v: ", i)
		fmt.Scanln(&tab1, &tab2, &tab3, &tab4)
		if (tab1 != "merah") || (tab2 != "kuning") || (tab3 != "hijau") || (tab4 != "ungu") {
			success = false
		}

	}
	fmt.Println("BERHASIL: ", success)
}
