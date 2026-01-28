package main

import "fmt"

func main() {
	var n int
	var arr [10][10]int
	fmt.Scan(&n)
	cnt := 1
	for i := range n {
		for j := range n {
			arr[j][i] = cnt
			cnt++
		}
	}	
	for i := range n {
		for j := range n {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}