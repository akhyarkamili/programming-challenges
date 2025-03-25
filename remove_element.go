package main

func removeElement(nums []int, val int) int {
	valIndices := make(chan int, len(nums))

	count := 0
	for i, v := range nums {
		if v == val {
			valIndices <- i
		} else {
			count++
			select {
			case valI := <-valIndices:
				swap(nums, valI, i)
				valIndices <- i
			default:
				break
			}
		}
	}

	return count
}
