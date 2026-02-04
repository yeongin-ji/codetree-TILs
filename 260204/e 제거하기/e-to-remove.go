package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	for i, v := range r {
		if v == 'e' {
			if i == len(r)-1 {
				r = r[:i]
			} else {
				r = append(r[:i], r[i+1:]...)
				break
			}
		}
	}	
	s = string(r)
	fmt.Println(s)
}