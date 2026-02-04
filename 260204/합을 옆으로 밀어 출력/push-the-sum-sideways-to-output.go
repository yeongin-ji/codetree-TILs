package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for range n {
		var v int
		fmt.Scan(&v)
		sum += v
	}	
	s := strconv.Itoa(sum)
	r := []rune(s)
	r = append(r[1:], r[0])
	s = string(r)
	fmt.Print(s)
}