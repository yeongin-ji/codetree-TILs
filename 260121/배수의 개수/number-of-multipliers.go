package main

import "fmt"

func main() {
	var cnt3, cnt5 int
	for range 10 {
		var n int
		fmt.Scanf("%d", &n)
		if n%3==0 {
			cnt3++
		}
		if n%5==0 {
			cnt5++
		}
	}	
	fmt.Print(cnt3, cnt5)
}