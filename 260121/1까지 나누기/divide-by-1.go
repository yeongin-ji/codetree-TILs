package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	i := 1
	cnt := 0
	for {
		n /= i
		cnt++
		if n <= 1 {
			break
		}
		i++
	}	
	fmt.Print(cnt)
}