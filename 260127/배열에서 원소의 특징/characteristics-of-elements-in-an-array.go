package main

import "fmt"

func main() {
	var arr [10]int
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	for i, v := range arr {
		if v%3==0 {
			fmt.Print(arr[i-1])
			break
		}
	}
}