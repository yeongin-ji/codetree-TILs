package main

import "fmt"

func main() {
	var n, m, cnt int
	fmt.Scan(&n, &m)
	var arr [110]int
	for i := range n {
		fmt.Scan(&arr[i])
	}	
	for i := range n {
		if arr[i]==m {
			cnt++
		}
	}
	fmt.Print(cnt)
}