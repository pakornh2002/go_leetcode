package minoperations

func MinOperations(logs []string) int {
	move := 0
	for _, str := range logs {
		if str == "./" || str == "../" && move == 0 {
			continue
		} else if str == "../" && move != 0 {
			move--
		} else {
			move++
		}
	}
	return move
}
