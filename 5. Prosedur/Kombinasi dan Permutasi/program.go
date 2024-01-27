package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var n_fact, nr_fact int
	factorial(n, &n_fact)
	factorial(n-r, &nr_fact)
	*hasil = n_fact / nr_fact
}

func combination(n, r int, hasil *int) {
	var n_fact, nr_fact, r_fact int

	factorial(n, &n_fact)
	factorial(r, &r_fact)
	factorial(n-r, &nr_fact)
	*hasil = n_fact / (r_fact * nr_fact)
}

func main() {
	var a, b, c, d, ac_comb, ac_perm, bd_comb, bd_perm int

	fmt.Scanln(&a, &b, &c, &d)

	permutation(a, c, &ac_perm)
	permutation(b, d, &bd_perm)
	combination(a, c, &ac_comb)
	combination(b, d, &bd_comb)
	fmt.Println(ac_perm, ac_comb)
	fmt.Println(bd_perm, bd_comb)
}
