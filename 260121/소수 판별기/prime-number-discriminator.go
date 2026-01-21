package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	flag := true
	for i := 2; i < n; i++ {
		if n%i==0 {
			flag = false
		}
	}	
	if flag {
		fmt.Print("P")
	} else {
		fmt.Print("C")
	}
}