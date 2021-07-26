/*
10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
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
	n := 10001
	val := 1
	count := 0

	for count <= n {
		for is_prime(val) == false {
			val++
		}
		val++
		count++
	}
	fmt.Println(val - 1)
}
