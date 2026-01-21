package main

import "fmt"

func main() {
	for {
		var n, cnt, sum int
		fmt.Scan(&n)
		if n < 20 || n >= 30 {
			break
		}
		sum += n
		cnt++
	}	
	avg := float64(sum)/float64(cnt)
	fmt.Printf("%.2f", avg)
}