package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	r1 := make([]rune, 0, 10)
	r2 := make([]rune, 0, 10)	
	for _, v := range s1 {
		if v>='0' && v<='9' {
			r1 = append(r1, v)
		}
	}
	for _, v := range s2 {
		if v>='0' && v<='9' {
			r2 = append(r2, v)
		}
	}
	s1 = string(r1)
	s2 = string(r2)
	a, _ := strconv.Atoi(s1)
	b, _ := strconv.Atoi(s2)
	fmt.Print(a+b)
}