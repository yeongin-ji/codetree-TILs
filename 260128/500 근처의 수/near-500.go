package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [20]int
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	
	max := math.MinInt
	min := math.MaxInt
	for i := range 10 {
		if arr[i] < 500 && arr[i] > max {
			max = arr[i]
		} else if arr[i] > 500 && arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Print(max, min)
}