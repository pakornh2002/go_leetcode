package rotate

import "fmt"

func Rotate(matrix [][]int) {
	queue := []int{}
	result := [][]int{}
	for j := 0; j < len(matrix); j++ {
		for k := len(matrix) - 1; k >= 0; k-- {
			queue = append(queue, matrix[k][j])
		}
		result = append(result, queue)
		queue = []int{}
	}
	copy(matrix, result)
	fmt.Println(matrix)
}
