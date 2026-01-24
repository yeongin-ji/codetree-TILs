package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n==1 {
		fmt.Print("*")
		return
	}
	// 높이는 짝수
	for i := range n-n%2 {
		for j := range n-n%2 {
			if i==0 {
				fmt.Print("* ")
			} else if j%2==0 {
				fmt.Print("  ")
			} else {
				if i <= j {
					fmt.Print("* ")
				} else {
					fmt.Print("  ")
				}
			}
		}
		if i==0 && n%2==1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}