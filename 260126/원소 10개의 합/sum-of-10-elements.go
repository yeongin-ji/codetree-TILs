package main

import "fmt"

const MAX = 10

func main() {
	var n [MAX]int
	for i := range MAX {
		fmt.Scan(&n[i])
	}	
	sum := 0
	for _, v := range n {
		sum += v
	}
	fmt.Print(sum)
}