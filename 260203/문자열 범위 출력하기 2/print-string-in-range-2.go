package main

import (
	"fmt"
)

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	for i := range n {
		idx := len(s)-1-i
		if idx >= 0 {
			fmt.Printf("%c", s[idx])
		}
	}
}