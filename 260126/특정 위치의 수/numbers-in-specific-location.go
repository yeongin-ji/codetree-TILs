package main

import "fmt"

func main() {
	var arr [10]int
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	sum := arr[2] + arr[4] + arr[9]
	fmt.Print(sum)
}