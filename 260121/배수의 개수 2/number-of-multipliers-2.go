package main

import "fmt"

func main() {
	cnt := 0
	for range 10 {
		var n int
		fmt.Scanf("%d", &n)
		if n%2==1 {
			cnt++
		}
	}	
	fmt.Print(cnt)
}