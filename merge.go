package main

// [1, 2, 2, 0, 0, 0] and [ 1, 3, 6]
// we get [1, 1, 2, 2, 3, 6]
// [1, 1, 2, 2, 3, 6]
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	pMain := m + n - 1

	for pMain >= 0 {
		// compare value at two pointers
		if p2 < 0 || (p1 >= 0 && nums1[p1] >= nums2[p2]) {
			// swap
			swap(nums1, p1, pMain)
			pMain--
			p1--
		} else {
			nums1[pMain] = nums2[p2]
			pMain--
			p2--
		}
	}

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
