package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [20]int
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	min := 200
	for i := range n-1 {
		if arr[i+1]-arr[i] < min {
			min = arr[i+1]-arr[i]
		}
	}
	fmt.Print(min)
}