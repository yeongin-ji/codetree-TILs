package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sum := 0
	for i := 1; i < n; i++ {
		if n%i==0 {
			sum += i
		}
	}	
	if sum == n {
		fmt.Print("P")
	} else {
		fmt.Print("N")
	}
}