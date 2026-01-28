package main

import "fmt"

func main() {
	var arr [100][100]int
	cnt := 1
	var n, m int
	fmt.Scan(&n, &m)
	for i := range n {
		for j := range m {
			arr[i][j] = cnt
			cnt++
		}
	}
	for i := range n {
		for j := range m {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}