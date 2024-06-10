package heightchecker

import "sort"

func HeightChecker(heights []int) int {
	sortedHeight := make([]int, len(heights))
	copy(sortedHeight, heights)
	sort.Ints(sortedHeight)
	var count int

	for i := 0; i < len(heights); i++ {
		if sortedHeight[i] != heights[i] {
			count++
		}
	}
	return count
}
