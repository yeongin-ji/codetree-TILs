package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var arr [100][100]int
	cnt := 0
	for i := range m {
		if i%2==0 {
			for j := range n {
				arr[j][i] = cnt
				cnt++
			}
		} else {
			for j := range n {
				arr[n-1-j][i] = cnt
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