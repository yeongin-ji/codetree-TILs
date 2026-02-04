package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	r1 := rune(s1[0])
	r2 := rune(s2[0])
	var sub rune
	if r1 <= r2 {
		sub = r2 - r1
	} else {
		sub = r1 - r2
	}
	fmt.Printf("%d %d", r1+r2, sub)
}