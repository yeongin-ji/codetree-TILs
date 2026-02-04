package main

import (
	"fmt"
)

func main() {
	var s string
	var n rune
	fmt.Scan(&s, &n)
	r := rune(s[0])
	fmt.Printf("%d %c", r, n)
}