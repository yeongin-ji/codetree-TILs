package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if a+b == b+a {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}