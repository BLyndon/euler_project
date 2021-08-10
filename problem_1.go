/*
Multiples of 3 or 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("")
	fmt.Println("Problem 1: Multiples of 3 or 5")
	fmt.Println("")

	sum := 0
	lim := 1000

	fmt.Println("Limit:", 1000)

	for i := 0; i < lim; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum:", sum)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
