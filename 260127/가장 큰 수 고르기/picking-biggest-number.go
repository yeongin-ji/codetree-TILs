package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [10]int
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	max := math.MaxInt
	for i := range 10 {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Print(max)
}