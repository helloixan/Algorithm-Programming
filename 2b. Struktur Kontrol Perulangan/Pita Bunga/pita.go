package main

import "fmt"

func main() {
	var (
		N           int
		pita, bunga string
	)

	/* v1 */
	// fmt.Print("N: ")
	// fmt.Scanln(&N)
	// pita = "Pita: "
	// for i := 1; i <= N; i++ {
	// 	fmt.Printf("Bunga %v: ", i)
	// 	fmt.Scanln(&bunga)
	// 	pita = pita + bunga + " - "
	// }
	// fmt.Println(pita)

	/* v2 */
	for selesai := false; !selesai; {
		fmt.Printf("Bunga %v: ", N+1)
		fmt.Scanln(&bunga)
		if bunga == "SELESAI" {
			selesai = true
		} else {
			pita = pita + bunga + " - "
			N++
		}
	}
	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", N)
}
