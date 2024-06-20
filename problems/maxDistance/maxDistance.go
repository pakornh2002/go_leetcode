package maxdistance

import "sort"

func MaxDistance(position []int, m int) int {
	sort.Ints(position)
	left, right := 1, position[len(position)-1]-position[0]
	result := 0

	for left <= right {
		mid := left + (right-left)/2
		count, lastPos := 1, position[0] // Place the first ball at the first position

		for i := 1; i < len(position); i++ {
			if position[i]-lastPos >= mid {
				count++
				lastPos = position[i]
				if count == m {
					break
				}
			}
		}

		if count >= m {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
