package main

import "fmt"

func main() {
	cnt := 0
	for cnt < 3 {
		var n int
		fmt.Scanf("%d", &n)
		if n%2==0 {
			fmt.Println(n/2)
			cnt++
		}
	}	
}