package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	r = append(r[1:], r[0])
	s = string(r)
	fmt.Println(s)	
}