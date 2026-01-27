package main

import "fmt"

func main() {
	var arr [4][4]int
	for i := range 4 {
		for j := range 4 {
			fmt.Scan(&arr[i][j])
		}
	}
	for i := range 4 {
		sum := 0
		for j := range 4 {
			sum += arr[i][j]
		}
		fmt.Println(sum)
	}
}