package algirithm

func GenerateMatrix(n int) [][]int {

	var res []int

	matrix := make([][]int, n)
	marked := make([][]bool, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		marked[i] = make([]bool, n)
	}

	count := n * n

	y := 0
	x := 0

	for len(res) < count {
		for x <= len(matrix[0])-1 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
				matrix[y][x] = len(res)
			}
			if x < len(matrix[0])-1 && !marked[y][x+1] {
				x++
			} else {
				break
			}
		}
		//fmt.Printf("1. x %d y %d\n", x, y)
		//fmt.Println(res)
		for y <= len(matrix)-1 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
				matrix[y][x] = len(res)
			}
			if y < len(matrix)-1 && !marked[y+1][x] {
				y++
			} else {
				break
			}
		}
		//fmt.Printf("2. x %d y %d\n", x, y)
		//fmt.Println(res)
		for x >= 0 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
				matrix[y][x] = len(res)
			}
			if x > 0 && !marked[y][x-1] {
				x--
			} else {
				break
			}
		}
		//fmt.Printf("3. x %d y %d\n", x, y)
		//fmt.Println(res)
		for y >= 0 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
				matrix[y][x] = len(res)
			}
			if y > 0 && !marked[y-1][x] {
				y--
			} else {
				break
			}
		}
		//fmt.Printf("4. x %d y %d\n", x, y)
		//fmt.Println(res)
	}

	//for i := 0; i < len(matrix); i++ {
	//	fmt.Println(matrix[i])
	//}

	return matrix
}
