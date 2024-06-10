package subarraysdivbyk

func SubarraysDivByK(nums []int, k int) int {
	remainderCount := make(map[int]int)
	remainderCount[0] = 1
	sum := 0
	count := 0

	for _, num := range nums {
		sum += num
		mod := (sum%k + k) % k

		if freq, exists := remainderCount[mod]; exists {
			count += freq
		}

		remainderCount[mod]++
	}

	return count
}

// func SubarraysDivByK(nums []int, k int) int {
// 	if len(nums) == 1 {
// 		if nums[0]%k == 0 {
// 			return 1
// 		} else {
// 			return 0
// 		}
// 	}
// 	sum := 0
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		sum += nums[i]
// 		mod := sum % k
// 		if mod == 0 {
// 			count++
// 		}
// 		for j := i + 1; j < len(nums); j++ {
// 			sum += nums[j]
// 			mod = sum % k
// 			if mod == 0 {
// 				count++
// 			}
// 		}
// 		sum = 0
// 	}
// 	return count
// }
