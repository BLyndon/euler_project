/*
Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("Problem 16: Power digit sum")
	fmt.Println("")

	var two = big.NewInt(2)
	var thousand = big.NewInt(1000)

	exp := two.Exp(two, thousand, nil)

	big_str := exp.String()

	sum := 0

	for i := 0; i < len(string(big_str)); i++ {
		val, _ := strconv.Atoi(string(big_str[i]))
		sum += val
	}

	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
