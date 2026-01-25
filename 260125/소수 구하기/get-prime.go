package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := range n {
		cnt := 0
		for j := range i+1 {
			if (i+1)%(j+1)==0 {
				cnt++
			}
		}
		if cnt==2 {
			fmt.Printf("%d ", i+1)
		}
	}	
}