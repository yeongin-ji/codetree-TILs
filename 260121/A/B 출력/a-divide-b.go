package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	
	rst := a / b
	rmd := a % b * 10
	fmt.Printf("%d.", rst)
	for range 20 {
		rst = rmd / b
		rmd = rmd % b * 10
		fmt.Print(rst)
	}
}