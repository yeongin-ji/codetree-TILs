package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var m, m2 int
	fmt.Scan(&n)
	var arr [110]int
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	m = arr[0]
	idx := -1
	for i := range n {
		if arr[i] > m {
			m = arr[i]
			idx = i
		}
	}
	arr[idx] = math.MinInt
	m2 = arr[0]
	for i := range n {
		if arr[i] > m2 {
			m2 = arr[i]
		}
	}
	fmt.Print(m, m2)
}