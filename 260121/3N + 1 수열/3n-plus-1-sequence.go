package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	for {
		if n == 1 {
			break
		}
		if n%2==0 {
			n /= 2
			cnt++
		} else {
			n = n*3 + 1
			cnt++
		}
	}
	fmt.Print(cnt)
}