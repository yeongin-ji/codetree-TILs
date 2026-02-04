package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	for _, r := range str {
		if r>='a' && r<='z' {
			fmt.Printf("%c", r-32)
		} else if r>='A' && r<='Z' {
			fmt.Printf("%c", r)
		}
	}
}