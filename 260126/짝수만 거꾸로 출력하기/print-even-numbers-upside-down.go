package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [100]int
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	for i := n-1; i >= 0; i-- {
		if arr[i]%2==0 {
			fmt.Printf("%d ", arr[i])
		}
	}
}