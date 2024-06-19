package mindays

import "math"

func MinDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	canMakeBouquets := func(days int) bool {
		bouquets := 0
		flowers := 0

		for _, bloom := range bloomDay {
			if bloom <= days {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			} else {
				flowers = 0
			}

			if bouquets >= m {
				return true
			}
		}

		return bouquets >= m
	}

	left, right := math.MaxInt32, math.MinInt32

	for _, day := range bloomDay {
		if day < left {
			left = day
		}
		if day > right {
			right = day
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if canMakeBouquets(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
