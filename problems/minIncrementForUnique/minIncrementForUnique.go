package minincrementforunique

import "sort"

func MinIncrementForUnique(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			increment := nums[i-1] - nums[i] + 1
			nums[i] += increment
			count += increment
		}
	}
	return count
}

// func MinIncrementForUnique(nums []int) int {
// 	checkDup := make(map[int]int)
// 	i := 0
// 	count := 0
// 	for i < len(nums) {
// 		checkDup[nums[i]]++
// 		if checkDup[nums[i]] > 1 {
// 			nums[i]++
// 			count++
// 		} else {
// 			i++
// 		}
// 	}
// 	return count
// }
