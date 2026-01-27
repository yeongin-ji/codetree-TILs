package main

import "fmt"

func main() {
	var arr [10]int	
	for {
		var v int
		if _, err := fmt.Scan(&v); err != nil && v==0 {
			break
		}
		arr[v/10]++
	}
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d - %d\n", i, arr[i])
	}
}