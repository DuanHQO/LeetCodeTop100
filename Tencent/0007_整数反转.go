package Tencent

func Reverse(x int) int {
	if x == 0 {
		return x
	}

	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}

	if ret > 1<<31 || ret < -1<<31 {
		return 0
	}
	return ret
}
