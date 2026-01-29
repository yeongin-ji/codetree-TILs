package main

import "fmt"

func main() {
	var n, m int
	var arr [10][10]int
	fmt.Scan(&n, &m)
	for range m {
		var r, c int
		fmt.Scan(&r, &c)
		arr[r][c] = 1
	}
	for i := range n {
		for j := range n {
			fmt.Printf("%d ", arr[i+1][j+1])
		}
		fmt.Println()
	}
}