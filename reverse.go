package main

// [a, b, c, d]
// [d, c, b, a] swap i with (len(s) - 1 - i) while i < mid
func reverseString(s []byte) {
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		if i < mid {
			s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
		}
	}
}
