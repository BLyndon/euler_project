package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("Problem 6: Sum square difference")
	fmt.Println("")

	num := 100
	sum := 0
	sum_sq := 0

	fmt.Println("n: ", num)

	for i := 1; i <= num; i++ {
		sum += i
		sum_sq += i * i
	}
	fmt.Println(sum*sum, "-", sum_sq, "=", sum*sum-sum_sq)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
