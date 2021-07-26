/*
Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"math"
)

func get_digit(num int, digit int) int {
	expo := int(math.Pow(10, float64(digit-1)))

	return (num / expo) % 10
}

func is_palin(n int) bool {
	digits := 0

	n_ := n
	for n_ > 0 {
		n_ /= 10
		digits += 1
	}

	for i := 1; i <= digits/2; i++ {
		dig1 := get_digit(n, i)
		dig2 := get_digit(n, digits-(i-1))

		if dig1 != dig2 {
			return false
		}
	}

	return true
}

func main() {
	var I, J int
	max_palin := 0

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			if is_palin(i * j) {
				if i*j > max_palin {
					max_palin = i * j
					I, J = i, j
				}
			}
		}
	}
	fmt.Println(I, "x", J, "=", I*J)
}
