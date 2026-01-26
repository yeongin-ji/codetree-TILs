package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [110]int
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	for i := range n {
		if arr[i]%2==0 {
			fmt.Printf("%d ", arr[i])
		}
	}
}