package piscine

func HashCode(dec string) string {
	res:=""
	for _,r:=range dec{
		result:=(int(r)+len(dec))%127
		if result<32 && result>126{
			result+=33
		}
		res+=string(rune(result))
	}
	return res
}