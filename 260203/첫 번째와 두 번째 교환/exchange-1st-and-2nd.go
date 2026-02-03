package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	t1 := r[0]
	t2 := r[1]
	for i, v := range r {
		if v == t1 {
			r[i] = t2
		}
		if v == t2 {
			r[i] = t1
		}
	}
	s = string(r)
	fmt.Print(s)
}