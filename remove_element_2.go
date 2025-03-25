package main

func removeElement2(nums []int, val int) int {
	nums2 := make([]int, len(nums))
	count := 0
	for i := range nums {
		if nums[i] != val {
			nums2[count] = nums[i]
			count++
		}
	}

	for i := 0; i < count; i++ {
		nums[i] = nums2[i]
	}

	return count
}
