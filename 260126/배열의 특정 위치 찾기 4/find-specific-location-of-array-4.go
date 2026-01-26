package main

import "fmt"

func main() {
	var arr [10]int
	var sum, cnt int
	for i := range 10 {
		fmt.Scan(&arr[i])
	}	
	for _, v := range arr {
		if v==0 {
			break
		}
		if v%2==0 {
			sum += v
			cnt++
		}
	}
	fmt.Print(cnt, sum)
}