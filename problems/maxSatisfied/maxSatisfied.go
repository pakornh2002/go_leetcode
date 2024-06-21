package maxsatisfied

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	totalSatisfied := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			totalSatisfied += customers[i]
		}
	}

	maxAdditionalSatisfied := 0
	currentAdditionalSatisfied := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			currentAdditionalSatisfied += customers[i]
		}
	}

	maxAdditionalSatisfied = currentAdditionalSatisfied
	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 1 {
			currentAdditionalSatisfied += customers[i]
		}
		if grumpy[i-minutes] == 1 {
			currentAdditionalSatisfied -= customers[i-minutes]
		}
		if currentAdditionalSatisfied > maxAdditionalSatisfied {
			maxAdditionalSatisfied = currentAdditionalSatisfied
		}
	}

	return totalSatisfied + maxAdditionalSatisfied
}
