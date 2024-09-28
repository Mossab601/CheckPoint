package main

import(
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AtoiBase(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1 
		}
		res = res*10 + int(s[i]-'0')
	}
	return res
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	nb:= AtoiBase(os.Args[1])
	if nb <= 0 {
		fmt.Println(0)
		return
	}
	sum := 0
	for i := 2; i <= nb; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
	fmt.Println(nb)
}


func atoit(s string) int{
	n:=0
	for i := 0; i < len(s); i++ {
		n=n*10+int(s[i]-'0')
	}
	return n
}