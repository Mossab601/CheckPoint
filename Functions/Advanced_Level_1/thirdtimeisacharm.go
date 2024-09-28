package piscine

func ThirdTimeIsACharm(str string) string {
	res := ""
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	return res + "\n"
}
