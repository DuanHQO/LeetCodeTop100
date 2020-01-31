package algirithm

func subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	var result [][]int

	n := len(nums)

	for i := 0; i <= n; i++ {
		for j := 0; j+i <= n; j++ {

			result = append(result)
		}
	}

	return result
}

func enumSubsets(nums []int, com int) {

}
