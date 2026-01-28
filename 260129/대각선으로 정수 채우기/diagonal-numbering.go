package main

import "fmt"

func main() {
	var n, m int
	var arr [100][100]int
	fmt.Scan(&n, &m)
	cnt := 1
	for i := range m {
		for j := range n {
			if i-j >= 0 {
				arr[j][i-j] = cnt
				cnt++
			}
		}
	}	
	for i := 1; i < n; i++ {
		for j := range n-1 {
			if i+j < n && m-1-j >= 0 {
				arr[i+j][m-1-j] = cnt
				cnt++
			}
		}
	}
	for i := range n {
		for j := range m {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}