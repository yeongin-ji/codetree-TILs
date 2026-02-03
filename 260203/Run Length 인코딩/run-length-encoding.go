package main

import (
	"fmt"
)

func main() {
	var s, rst string
	fmt.Scan(&s)
	cnt := 1
	for i := range len(s) {
		if i!=len(s)-1 && s[i]==s[i+1] {
			cnt++
		} else {
			rst += fmt.Sprintf("%c%d", s[i], cnt)
			cnt = 1
		}
	}	
	fmt.Println(len(rst))
	fmt.Println(rst)
}