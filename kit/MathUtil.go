package kit

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxPoint(a, b *int) int {
	if *a < *b {
		return *a
	} else {
		return *b
	}
}

func MinPoint(a, b *int) int {
	if *a < *b {
		return *a
	}
	return *b
}
