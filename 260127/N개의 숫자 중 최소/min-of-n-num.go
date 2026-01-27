package main

import (
	"fmt"
	"math"
)

const MAX = math.MaxInt

func main() {
	var n int
	var arr [110]int
	fmt.Scan(&n)
	for i := range n {
		fmt.Scan(&arr[i])
	}
	min := MAX
	for i := range n {
		if arr[i] < min {
			min = arr[i]
		}
	}
	cnt := 0
	for i := range n {
		if arr[i] == min {
			cnt++
		}
	}
	fmt.Print(min, cnt)
}