package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	res := ""
	if len(args) != 3 {
		return
	}
	arg1 := args[0]
	arg2 := args[1]
	arg3 := args[2]
	for _, r := range arg1 {
		if string(r) == arg2 {
			res += arg3
		} else {
			res += string(r)
		}
	}
	fmt.Println(res)
}
