package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [10][10]int
	cnt := 1
	flag := true
	for i := range n {
		if flag {
			for j := range n {
				arr[n-1-j][n-1-i] = cnt
				cnt++
			}
			flag = !flag
		} else {
			for j := range n {
				arr[j][n-1-i] = cnt
				cnt++
			}
			flag = !flag
		}
	}	
	for i := range n {
		for j := range n {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}