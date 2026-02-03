package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	var eeCnt, ebCnt int
	for i, r := range s {
		if i!=len(s)-1 && r=='e' {
			if rune(s[i+1]) == 'e' {
				eeCnt++
			} else if rune(s[i+1]) == 'b' {
				ebCnt++
			}
		}
	}	
	fmt.Print(eeCnt, ebCnt)
}