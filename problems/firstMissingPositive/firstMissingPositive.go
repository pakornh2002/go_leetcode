package firstMissingPositive

import "sort"

func FirstMissingPositive(nums []int) int {
	sort.Ints(nums)

	current := 1

	for _, num := range nums {
		if num < current {
			continue
		}

		if num == current {
			current++
		} else if num > current {
			break
		}
	}

	return current
}

// func FirstMissingPositive(nums []int) int {
// 	sort.Ints(nums)
// 	var current int
// 	for i, num := range nums {
// 		if num < 0 && len(nums) == 1 || i == 0 && num > 1 {
// 			return 1
// 		} else if num < 0 {
// 			continue
// 		} else if num-1 != current && i != 0 && num != current {
// 			if num-1 != current+1 {
// 				return current + 1
// 			} else {
// 				current = num - 1
// 				break
// 			}
// 		} else {
// 			current = num
// 		}
// 	}
// 	if current == nums[len(nums)-1] {
// 		current = nums[len(nums)-1] + 1
// 	} else if current == 0 {
// 		current = 1
// 	}
// 	return current
// }
