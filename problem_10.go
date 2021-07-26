/*
Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import "fmt"

func is_prime(val int) bool {
	for i := 2; i <= val/2; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	max := 2000000
	sum := 0
	for val := 2; val < max; val++ {
		if is_prime(val) {
			sum += val
		}
	}
	fmt.Println(sum)
}
