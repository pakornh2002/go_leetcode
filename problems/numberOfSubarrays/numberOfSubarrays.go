package numberofsubarrays

func NumberOfSubarrays(nums []int, k int) int {
	count := 0
	oddCount := 0
	prefixCount := make(map[int]int)
	prefixCount[0] = 1
	for _, num := range nums {
		if num%2 != 0 {
			oddCount++
		}

		if prefixCount[oddCount-k] > 0 {
			count += prefixCount[oddCount-k]
		}

		prefixCount[oddCount]++
	}

	return count
}
