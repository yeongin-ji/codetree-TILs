package main

import "fmt"

func main() {
	var arr1, arr2 [3][3]int
	for i := range 3 {
		for j := range 3 {
			fmt.Scan(&arr1[i][j])
		}
	}	
	for i := range 3 {
		for j := range 3 {
			fmt.Scan(&arr2[i][j])
		}
	}
	for i := range 3 {
		for j := range 3 {
			fmt.Printf("%d ", arr1[i][j]*arr2[i][j])
		}
		fmt.Println()
	}

}