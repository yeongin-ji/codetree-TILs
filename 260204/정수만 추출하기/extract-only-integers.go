package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	var idx1, idx2 int
	var a, b int
	flag := true
	aflag, bflag := false, false
	for i, v := range s1 {
		if v<'0' || v>'9' {
			idx1 = i
			flag = false
			break
		}
	}	
	if flag {
		a, _ = strconv.Atoi(s1)
		aflag = true
	}
	
	flag = true
	for i, v := range s2 {
		if v<'0' || v>'9' {
			idx2 = i
			flag = false
			break
		}
	}
	if flag {
		b, _ = strconv.Atoi(s2)
		bflag = true
	}

	if !aflag {
		b1 := s1[:idx1]
		s1 = string(b1)
		a, _ = strconv.Atoi(s1)
	}
	if !bflag {
		b2 := s2[:idx2]
		s2 = string(b2)
		b, _ = strconv.Atoi(s2)
	}
	fmt.Print(a+b)
}