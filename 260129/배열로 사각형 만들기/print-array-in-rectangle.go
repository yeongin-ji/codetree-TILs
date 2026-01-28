package main

import "fmt"

func main() {
	var arr [5][5]int
	for i := range 5 {
		arr[0][i] = 1
		arr[i][0] = 1
	}	
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			arr[i][j] = arr[i-1][j]+arr[i][j-1]
		}
	}
	for i := range 5 {
		for j := range 5 {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}