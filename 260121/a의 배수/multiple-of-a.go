package main

import "fmt"

func main() {
	var N, a int
	fmt.Scanf("%d %d", &N, &a)
	for i := 1; i <= N; i++ {
		if i%a==0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}	
}