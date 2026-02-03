package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	r[1]='a'
	r[len(r)-2]='a'
	for _, v := range r {
		fmt.Printf("%c", v)	
	}
}