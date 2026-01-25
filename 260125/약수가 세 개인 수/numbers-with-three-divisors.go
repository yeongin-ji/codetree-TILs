package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	rst := 0
	for i := a; i <= b; i++ {
		cnt := 0
		for j := range i {
			if i%(j+1)==0 {
				cnt++
			}
		}
		if cnt==3 {
			rst++
		}
	}	
	fmt.Print(rst)
}