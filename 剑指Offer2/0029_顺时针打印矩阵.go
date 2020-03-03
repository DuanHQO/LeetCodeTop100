package 剑指Offer2

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	var ret []int
	if len(matrix) == 0 {
		return ret
	}

	width := len(matrix[0]) - 1
	height := len(matrix) - 1

	left := 0
	right := width
	top := 0
	bottom := height

	y, x := 0, 0
	direction := 0

	for len(ret) < (width+1)*(height+1) {
		ret = append(ret, matrix[y][x])
		switch direction {
		case 0:
			if x >= right {
				direction = 1
				top++
				y++
				continue
			} else {
				x++
			}
		case 1:
			if y >= bottom {
				direction = 2
				right--
				x--
				continue
			} else {
				y++
			}
		case 2:
			if x <= left {
				direction = 3
				bottom--
				y--
				continue
			} else {
				x--
			}
		case 3:
			if y <= top {
				direction = 0
				left++
				x++
				continue
			} else {
				y--
			}
		}
	}

	return ret
}
