package main

import "fmt"

func main() {
	const thr int = 25
	for {
		var n int
		fmt.Scanf("%d", &n)
		if n < thr {
			fmt.Println("Higher")
		} else if n > thr {
			fmt.Println("Lower")
		} else {
			fmt.Println("Good")
			break
		}
	}
}	