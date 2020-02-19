package algirithm

import "sort"

func LeastInterval(tasks []byte, n int) int {
	if tasks == nil {
		return 0
	}

	if n == 0 {
		return len(tasks)
	}

	time := 0

	dic := make([]int, 26)

	for _, task := range tasks {
		dic[task-'A']++
	}
	sort.Ints(dic)

	for dic[25] > 0 {
		i := 0
		for i <= n {
			if dic[25] == 0 {
				break
			}
			if i < 26 && dic[25-i] > 0 {
				dic[25-i]--
			}
			time++
			i++
		}
		sort.Ints(dic)
	}

	return time
}
