package main

import "fmt"

func main() {
	var n int
	var arr [200]int
	fmt.Scan(&n)
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	for i := range n {
		fmt.Printf("%d ", arr[i]*arr[i])
	}
}