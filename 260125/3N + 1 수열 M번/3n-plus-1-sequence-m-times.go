package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	for range m {
		var n, cnt int
		fmt.Scan(&n)
		for n != 1 {
			if n%2==0 {
				n /= 2
			} else {
				n = n*3+1
			}
			cnt++
		}
		fmt.Println(cnt)
	}
}