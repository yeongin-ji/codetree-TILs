package main

import "fmt"

func main() {
	var a, b int
	var cnt [10]int
	fmt.Scan(&a, &b)
	for a > 1 {
		cnt[a%b]++
		a /= b
	}
	sum := 0
	for i := range b {
		sum += cnt[i] * cnt[i]
	}
	fmt.Print(sum)
}