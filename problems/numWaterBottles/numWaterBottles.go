package numwaterbottles

func NumWaterBottles(numBottles int, numExchange int) int {
	result := numBottles
	for numBottles >= numExchange {
		result += numBottles / numExchange
		numBottles = (numBottles / numExchange) + (numBottles % numExchange)
	}
	return result
}
