package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for range n {
		var a, b int
		fmt.Scan(&a, &b)
		sum := 0
		for i := a; i <= b; i++ {
			if i%2==0 {
				sum += i
			}
		}
		fmt.Println(sum)
	}	
}