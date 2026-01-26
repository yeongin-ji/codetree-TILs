package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	var sum float64
	for range n {
		var v float64
		fmt.Scan(&v)
		sum += v
	}	
	avg := sum/float64(cnt)
	fmt.Printf("%.1f\n", avg)
	if avg >= 4.0 {
		fmt.Println("Perfect")
	} else if avg >= 3.0 {
		fmt.Println("Good")
	} else {
		fmt.Println("Poor")
	}
}