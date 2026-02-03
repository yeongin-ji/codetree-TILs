package main

import (
	"fmt"
)

func main() {
	var n int
	var rst string
	fmt.Scan(&n)
	for range n {
		var s string
		fmt.Scan(&s)
		rst += s
	}	
	fmt.Println(rst)
}