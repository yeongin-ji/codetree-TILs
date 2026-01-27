package main

import "fmt"

func main() {
	var n int
	var cnt [10]int
	fmt.Scan(&n)
	for range n {
		var v int
		fmt.Scan(&v)
		cnt[v]++
	}
	for i := 1; i <= 9; i++ {
		fmt.Println(cnt[i])
	}
}