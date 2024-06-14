// TODO
package jump

func Jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++
			currentEnd = farthest

			if currentEnd >= n-1 {
				break
			}
		}
	}
	return jumps
}
