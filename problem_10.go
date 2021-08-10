/*
Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
	"math"
	"time"
)

func is_prime(val int) bool {
	sqrt := int(math.Sqrt(float64(val)))
	for i := 2; i <= sqrt; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("Problem 10: Summation of primes")
	fmt.Println("")

	max := 2000000
	sum := 0
	for val := 2; val < max; val++ {
		if is_prime(val) {
			sum += val
		}
	}
	fmt.Println("Sum:", sum)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
