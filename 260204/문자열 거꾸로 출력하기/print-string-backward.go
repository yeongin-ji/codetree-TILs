package main

import (
	"fmt"
)

func main() {
	for {
		var s string
		fmt.Scan(&s)
		if s == "END" {
			break
		}
		for i := range len(s) {
			fmt.Printf("%c", s[len(s)-1-i])
		}
		fmt.Println()
	}
}