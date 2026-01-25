package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	cnt := 0
	for i := a; i <= b; i++ {
		sum := 0
		for j := range i {
			if j!=0 && i%j==0 {
				sum += j
			}
		}
		if i==sum {
			cnt++
		}
	}	
	fmt.Print(cnt)
}