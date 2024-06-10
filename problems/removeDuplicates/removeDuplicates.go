package removeduplicates

func RemoveDuplicates(nums []int) int {
	counts := make(map[int]bool)
	result := []int{}
	for _, num := range nums {
		if !counts[num] {
			counts[num] = true
			result = append(result, num)
		}
	}
	copy(nums, result)
	nums = nums[:len(result)]
	return len(nums)
}

// func RemoveDuplicates(nums []int) int {
// 	counts := make(map[int]bool)
// 	newLength := 0
// 	for _, num := range nums {
// 		if !counts[num] {
// 			counts[num] = true
// 			nums[newLength] = num
// 			newLength++
// 		}
// 	}
// 	nums = nums[:newLength]
// 	return len(nums)
// }
