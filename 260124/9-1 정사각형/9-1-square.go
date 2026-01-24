package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 9
	for range n {
		for range n {
			fmt.Print(cnt)
			cnt = (cnt-2+9)%9+1
		}
		fmt.Println()
	}	
}