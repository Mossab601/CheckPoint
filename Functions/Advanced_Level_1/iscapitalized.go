package piscine

func IsCapitalized(s string) bool {
	if s ==""{
		return false
	}
	for i := 0; i < len(s); i++ {
		if i==0 && s[i]>='a' && s[i]<='z'{
			return false
		}
		if i<len(s)-1 && s[i]==' ' && s[i+1]>='a' && s[i+1]<='z'{
			return false
		}
	}
	return true
}

