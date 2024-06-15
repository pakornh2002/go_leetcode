package findmaximizedcapital

import "sort"

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make([][2]int, len(profits))
	for i := 0; i < len(profits); i++ {
		projects[i] = [2]int{capital[i], profits[i]}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i][0] < projects[j][0]
	})

	currentCapital := w
	projectIndex := 0
	profitHeap := []int{}

	for i := 0; i < k; i++ {
		for projectIndex < len(projects) && projects[projectIndex][0] <= currentCapital {
			profitHeap = append(profitHeap, projects[projectIndex][1])
			projectIndex++
		}

		sort.Slice(profitHeap, func(i, j int) bool {
			return profitHeap[i] > profitHeap[j]
		})

		if len(profitHeap) == 0 {
			break
		}

		currentCapital += profitHeap[0]
		profitHeap = profitHeap[1:]
	}

	return currentCapital
}
