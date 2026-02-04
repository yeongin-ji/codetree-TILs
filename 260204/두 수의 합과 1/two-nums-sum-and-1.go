package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := a + b
	s := strconv.Itoa(sum)
	cnt := 0
	for _, v := range s {
		if v == '1' {
			cnt++
		}
	}
	fmt.Print(cnt)
}