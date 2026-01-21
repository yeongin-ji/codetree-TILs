package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	flag := false
	for i := 2; i <= n-1; i++ {
		if n%i==0 {
			flag = true
		}
	}	
	if flag {
		fmt.Print("C")
	} else {
		fmt.Print("N")
	}
}