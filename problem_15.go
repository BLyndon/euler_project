/*
Lattice paths

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/

package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("Problem 15: Lattice paths")
	fmt.Println("")

	var L int64 = 20

	fmt.Println(L, "x", L, "grid")

	paths := new(big.Int)
	paths.Binomial(2*L, L)

	fmt.Println("#paths:", paths)

	elapsed := time.Since(start)
	fmt.Println("Execution time", elapsed)
	fmt.Println("")
}
