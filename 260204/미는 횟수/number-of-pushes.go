package main

import (
	"fmt"
)

func main() {
	var A, B string
	fmt.Scan(&A, &B)
	Ar := []rune(A)
	Br := []rune(B)
	flag := true
	for N := range len(A)-1 {
		if string(Ar)==string(Br) {
			fmt.Print(N)
			flag = false
			break
		}
		Ar = append(Ar[len(Ar)-1:], Ar[:len(Ar)-1]...)
	}	
	if flag {
		fmt.Print(-1)
	}
}