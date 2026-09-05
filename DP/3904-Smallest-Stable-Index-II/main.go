package main

func firstStableIndex(nums []int, k int) int {
	l := len(nums)
	minCheck := make([]int, l)
	maxCheck := make([]int, l)
	minNum := nums[l-1]
	maxNum := nums[0]

	for i := 0; i < l; i++ {
		minNum = min(minNum, nums[l-1-i])
		minCheck[l-1-i] = minNum

		maxNum = max(maxNum, nums[i])
		maxCheck[i] = maxNum
	}

	for i := 0; i < l; i++ {
		if maxCheck[i]-minCheck[i] <= k {
			return i
		}
	}

	return -1
}
