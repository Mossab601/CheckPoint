package piscine

func CheckNumber(arg string) bool {
	for i := 0; i < len(arg); i++ {
		if CheckRune(rune(arg[i])) == true {
			return true
		}
	}
	return false
}

func CheckRune(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
