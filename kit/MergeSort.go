package kit

func MergeSort(arr []int) []int {
	//if hi <= lo {
	//	return
	//}
	//
	//mid := lo + (hi - lo) >> 2
	//MergeSort(arr, lo, mid)
	//MergeSort(arr, mid+ 1, hi)
	//
	//merge(arr, lo, mid, hi)

	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := MergeSort(arr[0:i])
	right := MergeSort(arr[i:])

	result := merge(left, right)
	return result
}

func merge(left []int, right []int) (result []int) {

	m, n := 0, 0
	l, r := len(left), len(right)
	if r > 0 && l > 0 && left[l-1] > right[0] {
		for m < l && n < r {
			if left[m] < right[n] {
				result = append(result, left[m])
				m++
			} else {
				//这里可以统计逆序对， count += len(left) - m
				result = append(result, right[n])
				n++
			}
		}
	}
	result = append(result, left[m:]...)
	result = append(result, right[n:]...)

	return result
}

func mergeB(arr []int, lo, mid, hi int) {
	var temp []int
	m, n := 0, 0
	if mid > lo && hi > mid && arr[mid] > arr[mid+1] {
		for lo+m < mid+1 && mid+1+n < hi+1 {
			if arr[lo+m] < arr[mid+1+n] {
				temp = append(temp, arr[lo+m])
				m++
			} else {
				temp = append(temp, arr[mid+1+n])
				n++
			}
		}
	}
	temp = append(temp, arr[lo+m:mid+1]...)
	temp = append(temp, arr[mid+1+n:hi+1]...)
}

func mergeA(left, right []int) (result []int) {
	m, n := 0, 0
	l, r := len(left), len(right)
	if l > 0 && r > 0 && left[l-1] > right[0] {
		for m < l && n < r {
			if left[m] < right[n] {
				result = append(result, left[m])
				m++
			} else {
				result = append(result, right[n])
				n++
			}
		}
	}
	result = append(result, left[m:]...)
	result = append(result, right[n:]...)

	return result
}
