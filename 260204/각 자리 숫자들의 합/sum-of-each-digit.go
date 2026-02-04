package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := strconv.Itoa(n)
	var sum int
	for _, r := range s {
		sum += int(r-'0')
	}
	fmt.Print(sum)
}