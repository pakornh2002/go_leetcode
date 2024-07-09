package averagewaitingtime

func AverageWaitingTime(customers [][]int) float64 {
	time := 0
	totalWaitingTime := 0

	for i := 0; i < len(customers); i++ {
		arrivalTime := customers[i][0]
		cookingTime := customers[i][1]

		if time < arrivalTime {
			time = arrivalTime
		}

		time += cookingTime

		waitingTime := time - arrivalTime
		totalWaitingTime += waitingTime
	}

	return float64(totalWaitingTime) / float64(len(customers))
}
