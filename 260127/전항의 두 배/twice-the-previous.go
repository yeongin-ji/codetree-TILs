package main

import "fmt"

func main() {
	var arr [20]int
	fmt.Scan(&arr[0], &arr[1])
	i := 1	
	for range 8 {
		i++
		arr[i] = arr[i-1] + 2*arr[i-2]
	}
	for i := range i+1 {
		fmt.Printf("%d ", arr[i])
	}
}