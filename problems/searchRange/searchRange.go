package searchrange

func SearchRange(nums []int, target int) []int {
	findFirst := false
	findLast := false
	i := 0
	j := len(nums) - 1
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	for !(findFirst && findLast) {
		if nums[i] == target && !findFirst {
			findFirst = true
		} else if findFirst {
			findFirst = true
		} else {
			i++
		}

		if nums[j] == target && !findLast {
			findLast = true
		} else if findLast {
			findLast = true
		} else {
			j--
		}

		if i > len(nums)-1 || j < 0 {
			return []int{-1, -1}
		}
	}
	return []int{i, j}
}
