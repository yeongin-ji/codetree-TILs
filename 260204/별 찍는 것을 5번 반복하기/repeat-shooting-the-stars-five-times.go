package main

import (
	"fmt"
)

func f() {
	fmt.Printf("**********\n")
}

func main() {
	for range 5 {
		f()
	}
}