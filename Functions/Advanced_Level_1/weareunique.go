package piscine

func WeAreUnique(str1, str2 string) int {
	str := str1 + str2
	count := 0
	chars := map[string]int{}
	if str == "" {
		return -1
	}
	for _, s := range str {
		chars[string(s)]++
	}
	for _, r := range chars {
		if r == 1 {
			count++
		}
	}
	return count
}
