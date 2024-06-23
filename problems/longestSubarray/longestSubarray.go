package longestsubarray

func LongestSubarray(nums []int, limit int) int {
	minDeque := []int{}
	maxDeque := []int{}
	left := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] >= nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, right)

		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] <= nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, right)

		for nums[maxDeque[0]]-nums[minDeque[0]] > limit {
			left++
			if minDeque[0] < left {
				minDeque = minDeque[1:]
			}
			if maxDeque[0] < left {
				maxDeque = maxDeque[1:]
			}
		}

		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}
