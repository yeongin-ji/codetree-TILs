package main

import (
	"fmt"
)

func main() {
	var n int
	var A string
	fmt.Scan(&n, &A)
	cnt := 0
	for range n {
		var s string
		fmt.Scan(&s)
		if A == s {
			cnt++
		}
	}
	fmt.Print(cnt)
}