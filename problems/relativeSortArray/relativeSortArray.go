package relativesortarray

import "sort"

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	result := []int{}
	duplicated := make(map[int]bool)
	lastIndex := 0
	for _, num2 := range arr2 {
		for _, num1 := range arr1 {
			if num2 == num1 {
				duplicated[num1] = true
				result = append(result, num1)
				lastIndex++
			}
		}
	}

	sort.Ints(arr1)
	for _, num1 := range arr1 {
		if _, exist := duplicated[num1]; !exist {
			result = append(result, num1)
		}
	}
	return result
}
