package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	cnt := 0
	for i := 1; i <= n; i++ {
		if i%2!=0 && i%3!=0 && i%5!=0 {
			cnt++
		}
	}	
	fmt.Print(cnt)
}