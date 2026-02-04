package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	for _, v := range s {
		if v>='A' && v<='Z' {
			fmt.Printf("%c", v+32)
		} else if v>='a' && v<='z' {
			fmt.Printf("%c", v)
		} else if v>='0' && v<='9' {
			fmt.Printf("%c", v)
		}
	}	
}