package checkSubarraySum

func CheckSubarraySum(nums []int, k int) bool {
	var sum int
	isSubSum := true
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum%k == 0 {
				return isSubSum
			}
		}
		sum = 0
	}
	return !isSubSum
}

// func CheckSubarraySum(nums []int, k int) bool {
//     sum := 0
//     modMap := map[int]int{0: -1}

//     for i := 0; i < len(nums); i++ {
//         sum += nums[i]
//         mod := sum % k

//         if mod < 0 {
//             mod += k
//         }

//         if prevIndex, exists := modMap[mod]; exists {
//             if i - prevIndex > 1 {
//                 return true
//             }
//         } else {
//             modMap[mod] = i
//         }
//     }

//     return false
// }
