package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	for i := 1; i <= a; i++ {
		if (i%2!=0 || i%4==0) && (i/8)%2!=0 && i%7<4 {
			fmt.Printf("%d ", i)
		}
	}	
}