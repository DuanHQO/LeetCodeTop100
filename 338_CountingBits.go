package main

func countBits(num int) []int {
	result := make([]int, num+1)
	result[0] = 0

	for i := 1; i <= num; i++ {
		result[i] = result[i>>1] + i&1
	}

	return result
}
