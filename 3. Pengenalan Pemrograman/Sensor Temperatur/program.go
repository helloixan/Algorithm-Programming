package main

import "fmt"

func main() {
	var (
		temperature, min, max, sum, zeroflag, n int
		avg                                     float64
	)

	zeroflag, sum, min, max, n = 1, 0, 0, 0, 0
	for selesai := false; !selesai; {
		fmt.Scan(&temperature)
		sum += temperature
		n += 1

		if temperature > max {
			max = temperature
		} else if temperature < min {
			min = temperature
		}

		if (temperature == 0) && (zeroflag < 2) {
			zeroflag++
		} else if temperature == 0 {
			selesai = true
			n--
		} else if (zeroflag == 2) && (temperature != 0) {
			zeroflag--
		}
	}

	avg = float64(sum) / float64(n)
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(avg)
}
