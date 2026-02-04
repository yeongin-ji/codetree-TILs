package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	var idx1, idx2 int
	for i, v := range s1 {
		if v<'0' || v>'9' {
			idx1 = i
		}
	}	
	for i, v := range s2 {
		if v<'0' || v>'9' {
			idx2 = i
		}
	}
	b1 := s1[:idx1]
	b2 := s2[:idx2]
	s1 = string(b1)
	s2 = string(b2)
	a, _ := strconv.Atoi(s1)
	b, _ := strconv.Atoi(s2)
	fmt.Print(a+b)
}