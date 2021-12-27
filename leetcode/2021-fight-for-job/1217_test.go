package ojeveryday

// tag-[数学]
// leetcode50: Pow(x, n)
func myPow_(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	if n == 0 {
		return 1.0
	}
	y := myPow(x, n>>1)
	if n&1 == 1 {
		return x * y * y
	}
	return y * y
}
