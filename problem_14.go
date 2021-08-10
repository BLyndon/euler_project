/*
Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("Problem 14: Longest Collatz sequence")
	fmt.Println("")

	len := 0
	n_max := 1000000
	n_len_max := 0

	fmt.Println("Max starting point:", n_max)

	for n0 := 2; n0 < n_max; n0++ {
		n := n0
		count := 1
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			count++
		}
		if count > len {
			n_len_max = n0
			len = count
		}
	}
	fmt.Println("Starting point:", n_len_max)
	fmt.Println("Length:", len)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
