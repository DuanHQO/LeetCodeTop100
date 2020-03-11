package 每日1题_3月

func canThreePartsEqualSum(A []int) bool {
	if A == nil || len(A) == 0 {
		return false
	}

	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	if sum%3 != 0 {
		return false
	}

	avg := sum / 3

	count := 0
	sum = 0

	for i := 0; i < len(A); i++ {
		sum += A[i]
		if sum == avg {
			count++
			if count == 2 && i < len(A)-1 {
				return true
			}
			sum = 0
		}
	}

	return false
}
