package main

func shortestSequence(rolls []int, k int) int {
	//for i := 1; true; i++ {
	//	// disprove all permutations P(k,i) are subsequences of rolls
	//	disprove := disproveAllPermutationsAreSubsequences(rolls, k, i) // bound to be true before end of loop
	//	if disprove {
	//		return i
	//	}
	//
	//}
	return 0 // will never happen
}

// disprove there exists a permutation of P(k,i) that is subseq of rolls
func disproveAllPermutationsAreSubsequences(rolls []int, k, i int) bool {
	return false
}

// how to disprove?
// first way to disprove, sheer numbers.
// If length of rolls is less than P(k, i) == k^i, then khalas.
