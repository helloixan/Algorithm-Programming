package main

import "fmt"

func main(){
	var (
		nam float64
		nmk string
	)

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	switch {
	case nam > 80 && nam <= 100 :
		nmk = "A"
	case nam > 72.5 :
		nmk = "AB"
	case nam > 65 :
		nmk = "B"
	case nam > 57.5 :
		nmk = "BC"
	case nam > 50 :
		nmk = "C"
	case nam > 40 :
		nmk = "D"
	case nam <= 40 && nam >= 0:
		nmk = "E"
	default :
		nmk = "INVALID"
	}
	fmt.Println("Nilai Mata Kuliah: ", nmk)
}