package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	y := 0
	for i := 1; i <= n; i++ {
		if i%4!=0 {
			continue
		}
		if i%100==0 && i%400!=0 {
			continue
		}
		y++
	}	
	fmt.Print(y)
}