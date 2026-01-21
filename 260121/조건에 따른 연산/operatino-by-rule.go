package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	for n < 1000 {
		if n%2==0 {
			n = n*3 + 1
			cnt++
		} else {
			n = n*2 + 2
			cnt++
		}
	}
	fmt.Print(cnt)
}