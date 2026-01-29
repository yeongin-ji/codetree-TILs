package main

import "fmt"

func main() {
	var n int
	var arr [15][15]int
	fmt.Scan(&n)
	for i := range n {
		for j := range i+1 {
			if j!=0 && i!=j {
				arr[i][j] = arr[i-1][j-1]+arr[i-1][j]
			} else {
				arr[i][j] = 1
			}
		}
	}
	for i := range n {
		for j := range i+1 {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}