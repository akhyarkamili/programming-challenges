package main

func majorityElement(nums []int) []int {
	// at most there are three such elements
	// <-- push small numbers to the left
	// --> push big numbers to the right
	// observation: you can sort into three buckets in linear time (3n)
	// once we have partially sorted buckets, we can find the most frequent element in each bucket
	return nil
}

func majorityElementInBucket(nums []int) (el, count int) {
	// find the majority element in the bucket in linear time and O(1) space
	// if there is no majority element, return -1
	// if there is a majority element, return the majority element
	for _, num := range nums {
		if count == 0 {
			el = num
			count++
		} else if el == num {
			count++
		} else {
			count--
		}
	}

	// check if el is the majority element
	count = 0
	for _, num := range nums {
		if num == el {
			count++
		}
	}

	if count > len(nums)/2 {
		return el, count
	} else {
		return
	}
}
