package piscine

func Gcd(a, b uint) uint {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}