package main

func largestInteger(nums []int, k int) int {
	l := len(nums)
	res := -1
	check := make(map[int]int)
	for i := 0; i < l; i++ {
		check[nums[i]]++
	}

	if k == l {
		for i := 0; i < l; i++ {
			if res < nums[i] {
				res = nums[i]
			}
		}
	} else if k == 1 {
		for i := 0; i < l; i++ {
			if check[nums[i]] == 1 && res < nums[i] {
				res = nums[i]
			}
		}
	} else {
		if res < nums[0] && check[nums[0]] == 1 {
			res = nums[0]
		}
		if res < nums[l-1] && check[nums[l-1]] == 1 {
			res = nums[l-1]
		}
	}

	return res
}
