package main

import "fmt"

func main() {
	var n [10]int
	sum := 0
	cnt := 0
	for i := range 10 {
		fmt.Scan(&n[i])
		if n[i] < 250 {
			sum += n[i]
			cnt++
		} else {
			break
		}
	}
	avg := float64(sum)/float64(cnt)
	fmt.Printf("%d %.1f", sum, avg)
}