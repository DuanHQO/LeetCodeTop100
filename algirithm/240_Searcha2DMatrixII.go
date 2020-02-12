package algirithm

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	y := len(matrix) - 1
	x := 0

	for y >= 0 && x < len(matrix[0]) {

		if matrix[y][x] < target {
			x++
		} else if matrix[y][x] > target {
			y--
		} else {
			return true
		}

	}

	return false
}
