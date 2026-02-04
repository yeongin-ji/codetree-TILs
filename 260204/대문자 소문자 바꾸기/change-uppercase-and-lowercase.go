package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	for _, v := range s {
		if v>='a' && v<='z' {
			fmt.Printf("%c", v-32)
		} else {
			fmt.Printf("%c", v+32)
		}
	}	
}