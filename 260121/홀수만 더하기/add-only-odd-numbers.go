package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sum := 0
	for range n {
		var v int
		fmt.Scanf("%d", &v)
		if v%2==1 && v%3==0 {
			sum += v
		}
	}	
	fmt.Print(sum)
}