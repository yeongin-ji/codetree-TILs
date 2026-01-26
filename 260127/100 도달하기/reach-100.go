package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr [100]int
	arr[0] = 1
	arr[1] = n
	fmt.Printf("%d %d ", arr[0], arr[1])
	i := 2
	for {
		arr[i] = arr[i-2] + arr[i-1]
		fmt.Printf("%d ", arr[i])
		if arr[i] > 100 {
			break
		}
		i++
	}	
}