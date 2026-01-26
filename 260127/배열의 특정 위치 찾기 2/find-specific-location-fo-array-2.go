package main

import "fmt"

func main() {
	var arr [10]int	
	for i := range 10 {
		fmt.Scan(&arr[i])
	}
	var evensum, oddsum int
	for i, v := range arr {
		if (i+1)%2!=0 {
			oddsum += v	
		} else {
			evensum += v
		}
	}
	if oddsum <= evensum {
		fmt.Print(evensum-oddsum)
	} else {
		fmt.Print(oddsum-evensum)
	}
	
}