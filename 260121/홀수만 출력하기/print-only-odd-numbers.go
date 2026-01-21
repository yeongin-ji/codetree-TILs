package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for range n {
		var v int
		fmt.Scanf("%d", &v)
		if v%2==1 && v%3==0 {
			fmt.Println(v)
		}
	}	
}