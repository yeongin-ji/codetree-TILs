package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(s)
	r := []rune(s)
	for range len(r) {
		r = append(r[len(r)-1:], r[:len(r)-1]...)
		s = string(r)
		fmt.Println(s)
		r = []rune(s)
	}	
}