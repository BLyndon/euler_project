/*
Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"time"
)

func check_div(n int, count int) bool {
	for i := 1; i <= count; i++ {
		if n%i != 0 {
			return false
		}

	}

	return true
}

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("Problem 5: Smallest multiple")
	fmt.Println("")

	count := 20
	num := count

	fmt.Println("Numbers 1 -", num)

	for check_div(num, count) == false {
		num += 1
	}
	fmt.Println("Smallest multiple:", num)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
