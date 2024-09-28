package piscine

func DigitLen(n, base int) int {
	count := 0
	if base > 36 || base < 2 {
		return -1
	}
	if n < 0 {
		n *= -1
	}
	for n > 0 {
		count++
		n = n / base
	}
	return count
}
