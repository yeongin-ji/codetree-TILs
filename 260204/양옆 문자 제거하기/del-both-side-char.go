package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	r = append(r[:2], r[3:]...)
	r = append(r[:len(r)-2], r[len(r)-1])
	s = string(r)
	fmt.Println(s)	
}