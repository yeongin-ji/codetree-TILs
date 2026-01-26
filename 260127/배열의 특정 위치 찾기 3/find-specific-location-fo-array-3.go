package main

import "fmt"

func main() {
	var arr [100]int
	i := 0
	for {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			break
		}
		i++
	}	
	for i, v := range arr {
		if v==0 {
			fmt.Print(arr[i-3]+arr[i-2]+arr[i-1])
			break
		}
	}
}