package main

import (
	"fmt"
	"math"
	"sort"
)

func getMaximumSwap(n int) int {
	var digits []int

	// get the digits
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	if len(digits) == 0 {
		return n
	}

	// reverse digits
	reverse(digits)

	sorted := make([]int, len(digits))
	copy(sorted, digits)
	sort.Ints(sorted)
	reverse(sorted)

	// find leftmost digit
	leftMostIndex := 0
	for i := 0; i < len(digits); i++ {
		if digits[i] != sorted[i] {
			leftMostIndex = i
			break
		}
	}
	fmt.Printf("leftMostIndex: %d\n", leftMostIndex)

	// find biggest digits' index to the right most of leftmost
	biggestIndex := leftMostIndex
	for i := leftMostIndex + 1; i < len(digits); i++ {
		if digits[i] >= digits[biggestIndex] {
			biggestIndex = i
		}
	}

	// swap
	t := digits[leftMostIndex]
	digits[leftMostIndex] = digits[biggestIndex]
	digits[biggestIndex] = t

	// final number reconstruction
	f := 0
	for i := 0; i < len(digits); i++ {
		f += digits[i] * int(math.Pow10(len(digits)-1-i))
	}

	return f
}

func reverse(digits []int) {
	for i := 0; i < len(digits)/2; i++ {
		digits[i], digits[len(digits)-1-i] = digits[len(digits)-1-i], digits[i]
	}
}
