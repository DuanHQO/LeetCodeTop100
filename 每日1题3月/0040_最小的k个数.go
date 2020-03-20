package 每日1题3月

func GetLeastNumbers(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	var mergeSort func([]int, int, int)
	var merge func([]int, int, int, int)

	merge = func(arr []int, lo, mid, hi int) {
		var tmp []int
		m, n := 0, 0
		if mid >= lo && hi > mid && arr[mid] > arr[mid+1] {
			for lo+m <= mid && mid+1+n <= hi {
				if arr[lo+m] < arr[mid+1+n] {
					tmp = append(tmp, arr[lo+m])
					m++
				} else {
					tmp = append(tmp, arr[mid+1+n])
					n++
				}
			}
		}
		tmp = append(tmp, arr[lo+m:mid+1]...)
		tmp = append(tmp, arr[mid+1+n:hi+1]...)
		left := arr[:lo]
		right := arr[hi+1:]
		arr = append(append(left, tmp...), right...)
	}

	mergeSort = func(arr []int, lo, hi int) {
		if lo == hi {
			return
		}

		mid := lo + (hi-lo)/2
		mergeSort(arr, lo, mid)
		mergeSort(arr, mid+1, hi)
		merge(arr, lo, mid, hi)
	}

	mergeSort(arr, 0, len(arr)-1)

	return arr[:k]
}
