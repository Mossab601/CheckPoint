package piscine

func Itoa1(nbr int) string {
	res := ""
	for nbr > 0 {
		res = string(rune('0'+nbr%10)) + res
		nbr /= 10
	}
	return res
}

func ZipString(str string) string {
	count := 1
	res := ""
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 && str[i] == str[i+1] {
			count++
		} else {
			res += Itoa(count) + string(str[i])
			count = 1
		}
	}
	return res
}
