package main

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target {
		return letters[0]
	}

	right := len(letters) - 1
	left := 0
	ans := 0
	// iterations: in every iteration, we will make the search space smaller between left and right
	for left <= right {
		midpoint := (right + left) / 2
		if target < letters[midpoint] {
			ans = midpoint       // best answer we have right now
			right = midpoint - 1 // continue searching
		} else {
			left = midpoint + 1
		}
	}

	return letters[ans]
}
