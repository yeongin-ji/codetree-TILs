package main

import "fmt"

func main() {
	var cnt [10]int
	for range 10 {
		var v int
		fmt.Scan(&v)
		cnt[v]++
	}	
	for i := 1; i <= 6; i++ {
		fmt.Printf("%d - %d\n", i, cnt[i])
	}
}