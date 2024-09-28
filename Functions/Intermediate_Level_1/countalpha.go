package piscine

func CountAlpha(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if CheckAlpha(rune(s[i])) == true {
			count++
		}
	}
	return count
}

func CheckAlpha(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}
