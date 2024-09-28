package piscine

func RepeatAlpha(s string) string {
	res:=""
	for _,r:=range s{
		if r>='a' && r<='z'{
			countAlpha:=int(r-'a'+1)
			for i := 0; i < countAlpha; i++ {
				res+=string(r)
			}
		}else if r>='A' && r<='Z' {
			countAlpha:=int(r-'A'+1)
			for i := 0; i < countAlpha; i++ {
				res+=string(r)
			}
		}else {
			res+=string(r)
		}
	}
	return res
}
