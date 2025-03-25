package main

import (
	random "math/rand"
	"slices"
	"testing"
)

var rand *random.Rand = random.New(random.NewSource(1))

func TestLength(t *testing.T) {
	for i := 0; i < 100; i++ {
		nums, val := generateRandomInput(1 + rand.Intn(9))

		expectK := 0
		for _, n := range nums {
			if val != n {
				expectK += 1
			}
		}

		k := removeElement2(nums, val)

		if k != expectK {
			t.Errorf("k %d is not equal to expectK %d", k, expectK)
		}
	}
}

func TestElementRemoved(t *testing.T) {
	for i := 0; i < 100; i++ {
		nums, val := generateRandomInput(1 + rand.Intn(15))
		t.Logf("Nums case: %v\nVal:%d\n", nums, val)
		numsSorted := make([]int, len(nums))
		copy(numsSorted, nums)
		slices.Sort(numsSorted)
		var expectNums []int
		for _, v := range numsSorted {
			if v != val {
				expectNums = append(expectNums, v)
			}
		}

		k := removeElement2(nums, val)
		fail := false
		result := make([]int, 0, len(nums))
		for i := 0; i < k; i++ {
			result = append(result, nums[i])
		}

		slices.Sort(result)
		for j := 0; j < k; j++ {
			if result[j] != expectNums[j] {
				t.Errorf("Result mismatch\nVal:%d\nExpected:%v\nGot:%v\n", val, expectNums, nums)
				fail = true
			}
		}

		if !fail {
			t.Logf("Test case pass: %v, %d", nums, val)
		}

	}
}

func generateRandomInput(maxLen int) (nums []int, val int) {
	n := rand.Intn(maxLen)
	if n == 0 {
		n = 1
	}

	nums = make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(10)
	}

	iv := rand.Intn(n)
	return nums, nums[iv]
}
