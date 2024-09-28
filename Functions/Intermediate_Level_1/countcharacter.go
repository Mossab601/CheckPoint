package piscine

func CountChar(str string, c rune) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == byte(c) {
			count++
		}
	}
	return count
}
