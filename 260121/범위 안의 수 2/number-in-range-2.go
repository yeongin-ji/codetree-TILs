package main

import "fmt"

func main() {
	var sum, cnt int
	var avg float64
	for range 10 {
		var v int
		fmt.Scanf("%d", &v)
		if v >= 0 && v <= 200 {
			sum += v
			cnt++
		}
	}
	avg = float64(sum)/float64(cnt)
	fmt.Printf("%d %.1f", sum, avg)
}