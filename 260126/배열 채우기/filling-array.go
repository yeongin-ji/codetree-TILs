package main

import "fmt"

func main() {
	var a [10]int
	for i := range 10 {
		fmt.Scan(&a[i])
		if a[i]==0 {
			break
		}
	}
	for i := len(a)-1; i >= 0; i-- {
		if a[i]==0 {
			continue
		}
		fmt.Printf("%d ", a[i])
	}
}