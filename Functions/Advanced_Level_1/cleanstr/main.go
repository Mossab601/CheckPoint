package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println()
		return
	}
	arg1 := arg[0]
	if arg1 == "" {
		fmt.Println()
		return
	}
	sl := Split(arg1)
	res := sl[0]
	for i := 1; i < len(sl); i++ {
		res += " " + string(sl[i])
	}

	fmt.Println(res)
}

func Split(s string) []string {
	res := ""
	slice := []string{}
	for _, r := range s {
		if r == ' ' || r=='\t'{
			if res != "" {
				slice = append(slice, res)
				res = ""
			} else {
				continue
			}
		} else {
			res += string(r)
		}
	}
	if res != "" {
		slice = append(slice, res)
	}
	return slice
}
