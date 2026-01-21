package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		if i%2!=0 && i%10!=5 && (i%3!=0 || i%9==0) {
			fmt.Printf("%d ", i)
		}
	}	
}