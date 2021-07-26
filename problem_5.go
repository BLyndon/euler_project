/*
Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

func check_div(n int, count int) bool {
	for i := 1; i <= count; i++ {
		if n%i != 0 {
			return false
		}

	}

	return true
}

func main() {
	count := 20
	num := count

	for check_div(num, count) == false {
		num += 1
	}
	fmt.Println(num)

}
