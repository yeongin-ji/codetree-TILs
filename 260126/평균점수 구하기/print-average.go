package main

import "fmt"

func main() {
	var cnt int
	var sum float64
	for range 8 {
		var n float64
		fmt.Scan(&n)
		sum += n
		cnt++
	}	
	avg := sum/float64(cnt)
	fmt.Printf("%.1f", avg)
}