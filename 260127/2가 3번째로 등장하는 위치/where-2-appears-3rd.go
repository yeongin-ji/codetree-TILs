package main

import "fmt"

func main() {
	var n int
	var arr [110]int
	fmt.Scan(&n)
	for i := range n {
		fmt.Scan(&arr[i])
	}
	cnt2 := 0
	for i := range n {
		if arr[i]==2 {
			cnt2++
		}
		if cnt2==3 {
			fmt.Print(i+1)
			break
		}
	}
}