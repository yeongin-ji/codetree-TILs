package main

import "fmt"

func main() {
	var arr [4][4]int
	for i := range 4 {
		for j := range 4 {
			fmt.Scan(&arr[i][j])
		}
	}	
	sum := 0
	for i := range 4 {
		for j := range i+1 {
			sum += arr[i][j]
		}
	}
	fmt.Print(sum)
}