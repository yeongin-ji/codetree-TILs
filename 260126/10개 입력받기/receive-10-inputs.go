package main

import "fmt"

func main() {
	var sum, cnt int
	var arr [10]int
	for i := range 10 {
		fmt.Scan(&arr[i])
	}
	for i := range 10 {
		if arr[i] == 0 {
			break
		}
		sum += arr[i]
		cnt++
	}
	avg := float64(sum)/float64(cnt)
	fmt.Printf("%d %.1f", sum, avg)
}