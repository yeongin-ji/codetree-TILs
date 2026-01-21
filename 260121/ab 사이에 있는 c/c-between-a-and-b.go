package main

import "fmt"

func main() {
	var a, b, c int
	flag := false
	fmt.Scan(&a, &b, &c)
	for i := a; i <= b; i++ {
		if i%c==0 {
			flag = true
		}
	}
	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}