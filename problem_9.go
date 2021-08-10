/*
Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c2

For example, 3^2 + 42 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("Problem : Special Pythagorean triplet")
	fmt.Println("")

	val := 1000.0

	fmt.Println("Value triplet:", val)

	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			if float64(a+b)+math.Sqrt(float64(a*a+b*b)) == val {
				fmt.Println("Product:", float64(a*b)*math.Sqrt(float64(a*a+b*b)))
				break
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
