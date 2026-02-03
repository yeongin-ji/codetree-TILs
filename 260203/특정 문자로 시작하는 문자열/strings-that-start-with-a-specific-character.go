package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var arr [20]string
	for i := range n {
		fmt.Scan(&arr[i])
	}
	var c string
	fmt.Scan(&c)
	sum := 0
	cnt := 0
	for i := range n {
		if c[0] == arr[i][0] {
			sum += len(arr[i])
			cnt++
		}
	}
	fmt.Printf("%d %.2f", cnt, float64(sum)/float64(cnt))
}