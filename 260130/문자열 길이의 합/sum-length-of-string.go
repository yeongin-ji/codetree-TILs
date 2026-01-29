package main

import "fmt"

func main() {
	var n int
	var s [10]string
	fmt.Scan(&n)
	for i := range n {
		fmt.Scan(&s[i])
	}	
	aCnt := 0
	lCnt := 0
	for i := range n {
		v := s[i]
		lCnt += len(v)
		if v[0]=='a' {
			aCnt++
		}
	}
	fmt.Printf("%d %d", lCnt, aCnt)
}