package minmovestoseat

import "sort"

func MinMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	totalMoves := 0

	for i := range seats {
		if seats[i] > students[i] {
			totalMoves += seats[i] - students[i]
		} else {
			totalMoves += students[i] - seats[i]
		}
	}

	return totalMoves
}
