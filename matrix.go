package main

// matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// output: true

// [1, 2, 3, 4] search for 1.5
// indices: left := 0, right := 3 mid := 1
// move the right to mid; left := 0, right := 1, mid := 0
// move the right to mid again: left := 0, right = 0, mid := 0
func searchMatrix(matrix [][]int, target int) bool {
	// binary search which row
	left, right := 0, len(matrix)-1
	mid := len(matrix) / 2

	foundRow := false
	for left < right {
		firstNum := matrix[mid][0]
		lastNum := matrix[mid][len(matrix[mid])-1]
		if firstNum > target {
			right = mid
			mid = (left + right) / 2
		} else if firstNum <= target {
			if lastNum >= target {
				foundRow = true
				break
			} else {
				left = mid
				mid = (left + right) / 2
			}
		}
	}
	if matrix[mid][0] > target {
		return false
	}

	return foundRow
}
