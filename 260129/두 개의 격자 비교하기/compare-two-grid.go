package main

import "fmt"

func main() {
	var arr1, arr2 [10][10]int
	var n, m int
	fmt.Scan(&n, &m)
	for i := range n {
		for j := range m {
			fmt.Scan(&arr1[i][j])
		}
	}
	for i := range n {
		for j := range m {
			fmt.Scan(&arr2[i][j])
		}
	}
	for i := range n {
		for j := range m {
			if arr1[i][j]==arr2[i][j] {
				fmt.Printf("0 ")
			} else {
				fmt.Printf("1 ")
			}
		}
		fmt.Println()
	}
}