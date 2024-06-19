package threesumclosest

import "math"

func ThreeSumClosest(nums []int, target int) int {
	lenghtNums := len(nums)
	if lenghtNums == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	minRange := math.MaxInt32
	minSum := 0
	for i := 0; i < lenghtNums-2; i++ {
		for j := i + 1; j < lenghtNums-1; j++ {
			for k := j + 1; k < lenghtNums; k++ {
				if math.Abs(float64(target)-float64(nums[i]+nums[j]+nums[k])) < float64(minRange) {
					minRange = int(math.Abs(float64(target) - float64(nums[i]+nums[j]+nums[k])))
					minSum = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return minSum
}
