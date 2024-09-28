package piscine

func RetainFirstHalf(str string) string {
	if len(str)<2{
		return str
	}
	return str[:len(str)/2]
}