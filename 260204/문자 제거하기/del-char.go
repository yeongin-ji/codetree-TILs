package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	for len(r) > 1 {
		var n int
		fmt.Scan(&n)
		if n+1 < len(r) {
			r = append(r[:n], r[n+1:]...)
			s = string(r)
			fmt.Println(s)
			r = []rune(s)
		} else {
			r = r[:len(r)-1]
			s = string(r)
			fmt.Println(s)
			r = []rune(s)
		}
	}
}