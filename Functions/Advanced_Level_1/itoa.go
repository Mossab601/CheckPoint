package piscine


func Itoa(n int) string {
	result := ""
	isNigative:=false
	if n<0{
		n*=-1
		isNigative=true
	}
    if n == 0 {
        return "0"
    }
    for n > 0 {
        result = string(rune(n%10+'0')) + result
        n = n / 10
    }
	if isNigative{
		return "-"+result
	}
    return result
}