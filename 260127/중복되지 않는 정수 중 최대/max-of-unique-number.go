package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var arr [1100]int
	m := map[int]int{}
	for i := range n {
		fmt.Scan(&arr[i])
		m[arr[i]]++
	}
	max := math.MinInt
	for k, v := range m {
		if v == 1 && k > max {
			max = k
		}
	}
	if max != math.MinInt {
		fmt.Print(max)
	} else {
		fmt.Print(-1)
	}
}