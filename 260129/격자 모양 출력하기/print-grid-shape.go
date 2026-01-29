package main

import "fmt"

func main() {
	var n, m int
	var arr [20][20]int
	fmt.Scan(&n, &m)
	for range m {
		var r, c int
		fmt.Scan(&r, &c)
		arr[r][c] = r*c
	}	
	for i := range n {
		for j := range n {
			fmt.Printf("%d ", arr[i+1][j+1])
		}
		fmt.Println()
	}
}