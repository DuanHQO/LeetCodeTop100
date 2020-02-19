package algirithm

func setZeroes(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}

	var idx [][]int

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == 0 {
				idx = append(idx, []int{y, x})
			}
		}
	}

	for i := 0; i < len(idx); i++ {
		y := idx[i][0]
		x := idx[i][1]
		for tmpy := 0; tmpy < len(matrix); tmpy++ {
			if matrix[tmpy][x] != 0 {
				matrix[tmpy][x] = 0
			}
		}
		for tmpx := 0; tmpx < len(matrix[0]); tmpx++ {
			if matrix[y][tmpx] != 0 {
				matrix[y][tmpx] = 0
			}
		}
	}
}
