package main

import "fmt"

func main() {
	var n, sum, cnt int
	fmt.Scanf("%d", &n)
	for range n {
		var v int
		fmt.Scanf("%d", &v)
		sum += v
		cnt++
	}	
	avg := float64(sum)/float64(cnt)
	fmt.Printf("%d %.1f", sum, avg)
}