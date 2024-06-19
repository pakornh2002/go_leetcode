package threesum

import "sort"

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	if n < 3 {
		return result
	}

	sort.Ints(nums) // Sort nums array

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates of nums[i]
		}

		target := -nums[i]
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates of nums[left] and nums[right]
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// func ThreeSum(nums []int) [][]int {
// 	if len(nums) == 3 {
// 		if nums[0]+nums[1]+nums[2] != 0 {
// 			return [][]int{}
// 		} else {
// 			return [][]int{nums}
// 		}
// 	}

// 	result := [][]int{}
// 	isDuplicated := make(map[[3]int]bool)
// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i + 1; j < len(nums)-1; j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				queue := []int{}
// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					queue = append(queue, nums[i], nums[j], nums[k])
// 					sort.Ints(queue)
// 					if !isDuplicated[[3]int(queue)] {
// 						result = append(result, queue)
// 						isDuplicated[[3]int(queue)] = true
// 					}
// 				}
// 			}

// 		}
// 	}
// 	return result
// }
