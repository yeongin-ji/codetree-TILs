package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := rune(s[0])
	r = (r-1+26-'a')%26+'a'
	fmt.Printf("%c", r)
}