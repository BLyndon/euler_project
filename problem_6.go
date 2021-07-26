package main

import "fmt"

func main() {
	num := 100
	sum := 0
	sum_sq := 0

	for i := 1; i <= num; i++ {
		sum += i
		sum_sq += i * i
	}
	fmt.Println(sum*sum, "-", sum_sq, "=", sum*sum-sum_sq)
}
