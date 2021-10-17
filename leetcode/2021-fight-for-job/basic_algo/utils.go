package basic_algo

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinusAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
