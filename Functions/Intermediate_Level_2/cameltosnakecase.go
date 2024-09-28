package piscine

func CamelToSnakeCase(s string) string {
	if !IsCamelCase(s) {
		return s
	}
	var result []rune
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, r)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func IsCamelCase(s string) bool {
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			return false
		}
	}
	return true
}
