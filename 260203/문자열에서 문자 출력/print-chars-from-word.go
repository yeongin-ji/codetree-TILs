package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
}