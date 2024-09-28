package piscine


func FromTo(from int, to int) string {
	if from <0 || from>99 || to <0 || to>99{
		return"Invalid\n"
	}
	if from == to {
		return FormatNUmber(from)+"\n"
	}
	if from< to{
		return FormatNUmber(from)+", "+FromTo(from+1,to)
	}
	return FormatNUmber(from)+", "+FromTo(from-1,to)
}

func FormatNUmber(n int) string{
	if n<10{
		return "0"+string(rune('0'+n))
	}
	return string(rune('0'+(n/10)))+string(rune('0'+(n%10)))
}