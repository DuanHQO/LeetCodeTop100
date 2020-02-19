package algirithm

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}

	var res []int

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	marked := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		marked[i] = make([]bool, len(matrix[0]))
	}

	count := len(matrix) * len(matrix[0])

	y := 0
	x := 0

	for len(res) < count {
		for x <= len(matrix[0])-1 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
			}
			if x < len(matrix[0])-1 && !marked[y][x+1] {
				x++
			} else {
				break
			}
		}
		fmt.Printf("1. x %d y %d\n", x, y)
		fmt.Println(res)
		for y <= len(matrix)-1 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
			}
			if y < len(matrix)-1 && !marked[y+1][x] {
				y++
			} else {
				break
			}
		}
		fmt.Printf("2. x %d y %d\n", x, y)
		fmt.Println(res)
		for x >= 0 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
			}
			if x > 0 && !marked[y][x-1] {
				x--
			} else {
				break
			}
		}
		fmt.Printf("3. x %d y %d\n", x, y)
		fmt.Println(res)
		for y >= 0 {
			if !marked[y][x] {
				marked[y][x] = true
				res = append(res, matrix[y][x])
			}
			if y > 0 && !marked[y-1][x] {
				y--
			} else {
				break
			}
		}
		fmt.Printf("4. x %d y %d\n", x, y)
		fmt.Println(res)
	}

	//fmt.Println(res)

	return res
}
