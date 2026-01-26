package main

import "fmt"

func main() {
	var arr [20]int
	fmt.Scan(&arr[0], &arr[1])
	i := 2
	fmt.Printf("%d %d ", arr[0], arr[1])
	for range 8 {
		arr[i] = (arr[i-2]+arr[i-1])%10
		fmt.Printf("%d ", arr[i])
		i++
	}	
}