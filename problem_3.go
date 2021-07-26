/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func factor(pn, n int) (int, int) {
	if n%pn == 0 {
		return pn, n / pn
	} else {
		pn += 1
		return factor(pn, n)
	}
}

func main() {
	pn := 2
	num := 600851475143
	remainder := num

	for remainder != 1 {
		pn, remainder = factor(pn, remainder)
		fmt.Println(pn)
	}
	fmt.Println("Largest prime factor of", num, "is", pn)
}
