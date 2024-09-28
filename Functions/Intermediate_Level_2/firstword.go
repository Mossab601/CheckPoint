package piscine

func FirstWord(s string) string {
	tab := Split(s)
	if len(tab) == 0 {
		return ""+"\n"
	}
	return tab[0] + "\n"
}

func Split(s string) []string {
	tab := []string{}
	res := ""
	for _, r := range s {
		if r == ' ' {
			if res != "" {
				tab = append(tab, res)
				res = ""
			} else {
				continue
			}
		} else {
			res += string(r)
		}
	}
	if res != "" {
		tab = append(tab, res)
	}
	return tab
}
