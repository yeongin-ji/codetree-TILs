package main

import "fmt"

func main() {
	var n int
	var arr [10][10]int
	fmt.Scan(&n)
	for i := range n {
		for j := range n {
			if i!=0 && j!=0 {
				arr[i][j] = arr[i-1][j-1]+arr[i-1][j]+arr[i][j-1]
			} else {
				arr[i][j] = 1
			}
		}
	}	
	for i := range n {
		for j := range n {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}