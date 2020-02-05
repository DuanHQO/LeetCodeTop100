package algirithm

import (
	"fmt"
	"leetcodetop100/kit"
)

func SortColors(nums []int) {
	kit.MergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
