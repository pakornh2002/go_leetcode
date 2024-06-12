package sortcolors

import "fmt"

func SortColors(nums []int) {
	i := 0
	for i < len(nums) {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		i++
	}
	fmt.Println(nums)
}
