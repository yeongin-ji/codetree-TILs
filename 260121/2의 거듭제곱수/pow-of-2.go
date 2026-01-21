package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	for n != 1 {
		n /= 2
		cnt++
	}
	fmt.Print(cnt)
}