package findthewinner

func FindTheWinner(n int, k int) int {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}

	i := 0
	for len(arr) > 1 {
		i = (i + k - 1) % len(arr)
		arr = append(arr[:i], arr[i+1:]...)
	}
	return arr[0]
}
