package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	base := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	var mod int
	var div int
	for i, r := range arr {
		mod = int(r) % len(base)
		div = int(r) / len(base)
		z01.PrintRune(base[div])
		z01.PrintRune(base[mod])
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	result := ""
	for _, s := range arr {
		if s < 32 || s > 126 {
			result += "."
		} else {
			result += string(s)
		}
	}
	for _, res := range result {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
