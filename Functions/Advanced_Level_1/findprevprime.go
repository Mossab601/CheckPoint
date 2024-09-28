package piscine

func FindPrevPrime(nb int) int {
	for {
		if IsPrime(nb){
			return nb
		}
		nb--
	}
}


func IsPrime(nb int ) bool{
	count:=0
	if nb <=1{
		return false
	}
	for i:=2;i<=nb;i++{
		if nb%i==0{
			count++
		}
	}
	if count ==1{
		return true
	}
	return false
}