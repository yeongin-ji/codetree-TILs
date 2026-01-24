package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	for i := range n {
		if i%2==0 {
			for range n {
				cnt++
				fmt.Printf("%d ", cnt)
			}
		} else {
			for range n {
				cnt += 2
				fmt.Printf("%d ", cnt)
			}
		}
		fmt.Println()
	}	
}