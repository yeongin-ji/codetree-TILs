package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scanf("%d %d", &a, &b)
	for i := a; i <= b; i++ {
		if i%6==0 && i%8!=0 {
			sum += i
		}
	}	
	fmt.Print(sum)
}