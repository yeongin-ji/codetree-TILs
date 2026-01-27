package main

import "fmt"

func main() {
	var arr [110]int
	var min, max int
	i := 0
	for {
		if _, err := fmt.Scan(&arr[i]); err != nil || arr[i]==999 || arr[i]==-999 {
			break
		}
		i++
	}	
	min = arr[0]
	max = arr[0]
	for j := range i {
		if arr[j] > max {
			max = arr[j]
		}
		if arr[j] < min {
			min = arr[j]
		}
	}
	fmt.Print(max, min)
}