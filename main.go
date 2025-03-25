package main

import (
	"fmt"
)

func main() {
	var n int
	n = getMaximumSwap(2376)
	fmt.Printf("original: %d, n: %d\n", 2376, n)
	n = getMaximumSwap(1993)
	fmt.Printf("original: %d, n: %d\n", 1993, n)
	n = getMaximumSwap(3112)
	fmt.Printf("original: %d, n: %d\n", 3112, n)
	n = getMaximumSwap(0)
	fmt.Printf("original: %d, n: %d\n", 0, n)
	n = getMaximumSwap(1111)
	fmt.Printf("original: %d, n: %d\n", 1111, n)
	n = getMaximumSwap(9937)
	fmt.Printf("original: %d, n: %d\n", 9937, n)
}
