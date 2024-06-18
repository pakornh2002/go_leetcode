package maxprofitassignment

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	result := 0
	for i := 0; i < len(worker); i++ {
		maxProfit := 0
		for j := 0; j < len(difficulty); j++ {
			if worker[i] >= difficulty[j] {
				if profit[j] > maxProfit {
					maxProfit = profit[j]
				}
			}
		}
		result += maxProfit
	}
	return result
}
