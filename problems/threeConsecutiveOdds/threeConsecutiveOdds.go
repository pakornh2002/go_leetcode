package threeconsecutiveodds

func ThreeConsecutiveOdds(arr []int) bool {
	i := 0
	count := 0
	for i < len(arr) {
		if arr[i]%2 != 0 {
			count++
		} else {
			count = 0
		}

		if count == 3 {
			return true
		}
		i++
	}
	return false
}
